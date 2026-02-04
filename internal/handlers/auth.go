package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// ContextKey é um tipo para chaves de contexto
type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
)

// CustomClaims define as claims customizadas do JWT
type CustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// publicRoutes é a lista de rotas que não exigem autenticação
var publicRoutes = map[string]bool{
	"/api/v1/users/login": true,
}

// isPublicRoute verifica se uma rota é pública
func isPublicRoute(path string) bool {
	return publicRoutes[path]
}

// JWTMiddleware valida o token JWT e injeta o user ID no contexto
// Rotas públicas definidas em publicRoutes não exigem autenticação
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verifica se a rota é pública
		if isPublicRoute(r.URL.Path) {
			// Permite a requisição sem autenticação
			next.ServeHTTP(w, r)
			return
		}

		// Extrai o token do header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeErrorResponse(w, "Token de autorização não fornecido", http.StatusUnauthorized)
			return
		}

		// Remove o prefixo "Bearer " do token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// Significa que não tinha o prefixo "Bearer "
			writeErrorResponse(w, "Formato do token inválido. Use: Bearer <token>", http.StatusUnauthorized)
			return
		}

		// Parse e valida o token
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Valida o método de assinatura
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			writeErrorResponse(w, fmt.Sprintf("Token inválido: %v", err), http.StatusUnauthorized)
			return
		}

		// Extrai as claims do token
		claims, ok := token.Claims.(*CustomClaims)
		if !ok || !token.Valid {
			writeErrorResponse(w, "Token inválido ou expirado", http.StatusUnauthorized)
			return
		}

		// Valida se o user_id existe nas claims
		if claims.UserID == "" {
			writeErrorResponse(w, "Token não contém user_id válido", http.StatusUnauthorized)
			return
		}

		// Injeta o user ID no contexto
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)

		// Continua com a requisição
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext extrai o user ID do contexto da requisição
func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok || userID == "" {
		return uuid.Nil, fmt.Errorf("user_id não encontrado no contexto")
	}
	return uuid.MustParse(userID), nil
}

// writeErrorResponse escreve uma resposta de erro em JSON
func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}
