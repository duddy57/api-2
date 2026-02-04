package repository

import (
	"context"
	"errors"
	"fmt"
	"olidesk-api-2/internal/domains"
	"olidesk-api-2/internal/store/pgstore"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresUsersRepository struct {
	db   *pgstore.Queries
	pool *pgxpool.Pool
}

func NewPostgresUsersRepository(db *pgxpool.Pool) UserRepository {
	return &postgresUsersRepository{db: pgstore.New(db), pool: db}
}

func (p *postgresUsersRepository) Save(users *domains.User, ctx context.Context) (uuid.UUID, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to begin tx for RegisterTeam: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	qtx := p.db.WithTx(tx)

	result, err := qtx.CreateUserQuery(ctx, pgstore.CreateUserQueryParams{
		Email:        users.Email,
		PasswordHash: users.Password,
		Username:     users.Name,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, errors.New("email or username already exists")
		}
		return uuid.UUID{}, err
	}

	if err := qtx.CreateMemberQuery(ctx, pgstore.CreateMemberQueryParams{
		UserID: result,
		Role:   pgstore.MemberRole(users.Role),
	}); err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}

	return result, nil
}
func (p *postgresUsersRepository) FindByID(id uuid.UUID, ctx context.Context) (*domains.User, error) {
	user, err := p.db.GetUserByIdQuery(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &domains.User{
		ID:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt.UTC(),
		UpdatedAt: user.UpdatedAt.UTC(),
	}, nil
}
func (p *postgresUsersRepository) FindByEmail(email string, ctx context.Context) (*domains.User, error) {
	user, err := p.db.GetUserByEmailQuery(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &domains.User{
		ID:       user.ID,
		Name:     user.Username,
		Email:    user.Email,
		Password: user.PasswordHash,
	}, nil
}
func (p *postgresUsersRepository) Update(users *domains.User, ctx context.Context) error {
	if err := p.db.UpdateUserQuery(ctx, pgstore.UpdateUserQueryParams{
		Username: users.Name,
		ID:       users.ID,
	}); err != nil {
		return err
	}
	return nil
}
func (p *postgresUsersRepository) Delete(id uuid.UUID, ctx context.Context) error {
	if err := p.db.DeleteUserQuery(ctx, id); err != nil {
		return err
	}
	return nil
}
func (p *postgresUsersRepository) GetMembers(ctx context.Context) ([]*domains.Member, error) {
	members, err := p.db.GetMemberQuery(ctx)
	if err != nil {
		return nil, err
	}

	membersList := make([]*domains.Member, 0, len(members))
	for _, member := range members {
		membersList = append(membersList, &domains.Member{
			ID:    member.ID,
			Name:  member.Username,
			Email: member.Email,
			Role:  string(member.Role),
		})
	}

	return membersList, nil
}
