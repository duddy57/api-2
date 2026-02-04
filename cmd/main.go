package main

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"
	"olidesk-api-2/internal/handlers"
	"olidesk-api-2/internal/handlers/spec"
	"olidesk-api-2/internal/repository"
	"olidesk-api-2/internal/usecase"
	"olidesk-api-2/internal/utils/config"
	"olidesk-api-2/internal/utils/logger"
	"os/signal"
	"syscall"
	"time"

	"os"

	scalargo "github.com/bdpiprava/scalar-go"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/phenpessoa/gutils/netutils/httputils"
	"github.com/resend/resend-go/v3"
)

func init() {

}

func main() {
	gob.Register(uuid.UUID{})

	if os.Getenv("ENVIRONMENT") != "production" {
		_ = godotenv.Load()
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context) (err error) {
	r := chi.NewMux()
	cfg := config.Load()

	l, err := logger.NewLogger(cfg.Environment)
	if err != nil {
		return err
	}

	l = l.Named("journey_logger")
	defer func() { _ = l.Sync() }()

	mailer := resend.NewClient(cfg.ResendAPIKey)

	protectedRouter := chi.NewRouter()
	protectedRouter.Use(handlers.JWTMiddleware)
	r.Use(middleware.RequestID, middleware.Recoverer, httputils.ChiLogger(l))

	pool, err := pgxpool.New(ctx, cfg.GetDatabaseURL())
	if err != nil {
		return err
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	ur := repository.NewPostgresUsersRepository(pool)
	cr := repository.NewPostgresClientsRepository(pool)
	fr := repository.NewPostgresFormRepository(pool)

	us := usecase.NewUserService(ur, l, mailer)
	cs := usecase.NewClientService(cr, l)
	fs := usecase.NewFormService(fr, l)

	si := handlers.NewHandlers(l, us, cs, fs)

	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		// Generate beautiful docs from your OpenAPI spec
		html, err := scalargo.NewV2(
			scalargo.WithSpecDir("./internal/handlers/spec"), // or WithSpecURL, WithSpecBytes
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprint(w, html)
	})
	r.Mount("/api", spec.Handler(&si, spec.WithRouter(protectedRouter)))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	defer func() {
		const timeout = 30 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			fmt.Errorf(err.Error())
		}
	}()

	errChan := make(chan error, 1)

	go func() {
		fmt.Println("server starting on:", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		return nil
	case err := <-errChan:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}
