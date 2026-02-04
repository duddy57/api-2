package usecase

import (
	"context"
	"errors"
	"olidesk-api-2/internal/domains"
	"olidesk-api-2/internal/repository"
	"olidesk-api-2/internal/utils/tokens"

	"github.com/google/uuid"
	"github.com/resend/resend-go/v3"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	CreateUser(CreateUserInput, context.Context) (uuid.UUID, error)
	GetUser(uuid.UUID, context.Context) (*domains.User, error)
	UpdateUser(uuid.UUID, UpdateUserInput, context.Context) error
	DeleteUser(uuid.UUID, context.Context) error
	LoginUser(LoginUserInput, context.Context) (LoginUserOutput, error)
	GetMembers(ctx context.Context) ([]*domains.Member, error)
}

type userService struct {
	repo   repository.UserRepository
	logger *zap.Logger
	mail   *resend.Client
}

func NewUserService(repo repository.UserRepository, logger *zap.Logger, mail *resend.Client) UserUseCase {
	return &userService{repo: repo, logger: logger, mail: mail}
}

func (u *userService) CreateUser(p CreateUserInput, ctx context.Context) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("failed to generate password hash", zap.Error(err))
		return uuid.UUID{}, err

	}

	id, err := u.repo.Save(&domains.User{
		Name:     p.Name,
		Email:    p.Email,
		Password: hash,
		Role:     p.Role,
	}, ctx)
	if err != nil {
		u.logger.Error("failed to save user", zap.Error(err))
		return uuid.UUID{}, err
	}

	// go func() {
	// 	if _, err := u.mail.Emails.Send(&resend.SendEmailRequest{
	// 		To:      []string{p.Email},
	// 		Subject: "Welcome to Sperium",
	// 		Template: &resend.EmailTemplate{
	// 			Id: "order-confirmation",
	// 			Variables: map[string]interface{}{
	// 				"name":     p.Name,
	// 				"password": p.Password,
	// 			},
	// 		},
	// 	}); err != nil {
	// 		u.logger.Error("failed to send welcome email", zap.Error(err))
	// 	}
	// }()

	return id, nil
}

func (u *userService) DeleteUser(id uuid.UUID, ctx context.Context) error {
	if err := u.repo.Delete(id, ctx); err != nil {
		u.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	return nil
}

func (u *userService) GetMembers(ctx context.Context) ([]*domains.Member, error) {
	members, err := u.repo.GetMembers(ctx)
	if err != nil {
		u.logger.Error("failed to get members", zap.Error(err))
		return nil, err
	}
	return members, nil
}

func (u *userService) GetUser(id uuid.UUID, ctx context.Context) (*domains.User, error) {
	user, err := u.repo.FindByID(id, ctx)
	if err != nil {
		u.logger.Error("failed to get user", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *userService) LoginUser(p LoginUserInput, ctx context.Context) (LoginUserOutput, error) {
	user, err := u.repo.FindByEmail(p.Email, ctx)
	if err != nil {
		u.logger.Error("failed to get user", zap.Error(err))
		return LoginUserOutput{}, err
	}

	if !checkPassword(user.Password, p.Password) {
		u.logger.Error("invalid email or password")
		return LoginUserOutput{}, errors.New("invalid email or password")
	}

	token, err := tokens.GenerateJWT(user.ID.String(), user.Email)
	if err != nil {
		u.logger.Error("failed to generate token", zap.Error(err))
		return LoginUserOutput{}, err
	}

	return LoginUserOutput{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   tokens.GetTokenExpirationTime(),
	}, nil
}

func (u *userService) UpdateUser(id uuid.UUID, p UpdateUserInput, ctx context.Context) error {
	user, err := u.repo.FindByID(id, ctx)
	if err != nil {
		u.logger.Error("failed to get user", zap.Error(err))
		return err
	}
	if p.Name != nil {
		user.Name = *p.Name
	}
	if err := u.repo.Update(&domains.User{
		ID:   id,
		Name: user.Name,
	}, ctx); err != nil {
		u.logger.Error("failed to update user", zap.Error(err))
		return err
	}
	return nil
}

func checkPassword(hashedPassword []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil
}
