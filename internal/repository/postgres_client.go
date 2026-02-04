package repository

import (
	"context"
	"olidesk-api-2/internal/domains"
	"olidesk-api-2/internal/store/pgstore"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresClientsRepository struct {
	db   *pgstore.Queries
	pool *pgxpool.Pool
}

func NewPostgresClientsRepository(db *pgxpool.Pool) ClientRepository {
	return &postgresClientsRepository{db: pgstore.New(db), pool: db}
}

func (u *postgresClientsRepository) SaveClient(c *domains.Client, ctx context.Context) (uuid.UUID, error) {
	arg := pgstore.CreateClientQueryParams{
		Name: c.ClientName,

		Email:       pgtype.Text{String: c.Contact.Email, Valid: true},
		Phone:       pgtype.Text{String: c.Contact.Phone, Valid: true},
		ContactName: pgtype.Text{String: c.Contact.ResposableName, Valid: true},

		CnpjCpf:    pgtype.Text{String: c.CnpjOrCpf, Valid: true},
		ClientType: pgstore.ClientType(c.ClientType),

		PostalCode:   pgtype.Text{String: c.Address.PostalCode, Valid: true},
		Neighborhood: pgtype.Text{String: c.Address.Neighborhood, Valid: true},
		Country:      pgtype.Text{String: c.Address.Country, Valid: true},
		State:        pgtype.Text{String: c.Address.State, Valid: true},
		City:         pgtype.Text{String: c.Address.City, Valid: true},
		Street:       pgtype.Text{String: c.Address.Street, Valid: true},
		Number:       pgtype.Text{String: c.Address.Number, Valid: true},
		Complement:   pgtype.Text{String: c.Address.Complement, Valid: true},

		Latitude:  pgtype.Float8{Float64: c.Address.Latitude, Valid: true},
		Longitude: pgtype.Float8{Float64: c.Address.Longitude, Valid: true},
	}

	id, err := u.db.CreateClientQuery(ctx, arg)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
func (u *postgresClientsRepository) FindClientByID(id uuid.UUID, ctx context.Context) (*domains.Client, error) {
	client, err := u.db.GetClientByIdQuery(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domains.Client{
		ID:         client.ID,
		CnpjOrCpf:  client.CnpjCpf.String,
		ClientType: string(client.ClientType),
		ClientName: client.Name,
		Contact: domains.ContactPerson{
			Email:          client.Email.String,
			Phone:          client.Phone.String,
			ResposableName: client.ContactName.String,
		},
		Address: domains.Address{
			PostalCode:   client.PostalCode.String,
			Country:      client.Country.String,
			Neighborhood: client.Neighborhood.String,
			State:        client.State.String,
			City:         client.City.String,
			Street:       client.Street.String,
			Number:       client.Number.String,
			Complement:   client.Complement.String,
			Latitude:     client.Latitude.Float64,
			Longitude:    client.Longitude.Float64,
		},
		CreatedAt: client.CreatedAt.UTC(),
		UpdatedAt: client.UpdatedAt.UTC(),
	}, nil
}
func (u *postgresClientsRepository) ListClients(ctx context.Context) ([]*domains.Client, error) {
	cData, err := u.db.GetAllClientsQuery(ctx)
	if err != nil {
		return nil, err
	}
	clients := make([]*domains.Client, 0, len(cData))
	for _, client := range cData {
		clients = append(clients, &domains.Client{
			ID:         client.ID,
			CnpjOrCpf:  client.CnpjCpf.String,
			ClientType: string(client.ClientType),
			ClientName: client.Name,
			Contact: domains.ContactPerson{
				Email:          client.Email.String,
				Phone:          client.Phone.String,
				ResposableName: client.ContactName.String,
			},
			Address: domains.Address{
				PostalCode:   client.PostalCode.String,
				Country:      client.Country.String,
				Neighborhood: client.Neighborhood.String,
				State:        client.State.String,
				City:         client.City.String,
				Street:       client.Street.String,
				Number:       client.Number.String,
				Complement:   client.Complement.String,
				Latitude:     client.Latitude.Float64,
				Longitude:    client.Longitude.Float64,
			},
			CreatedAt: client.CreatedAt.UTC(),
			UpdatedAt: client.UpdatedAt.UTC(),
		})
	}

	return clients, nil
}
func (u *postgresClientsRepository) UpdateClient(c *domains.Client, ctx context.Context) error {
	arg := pgstore.UpdateClientQueryParams{
		ID:   c.ID,
		Name: c.ClientName,

		Email:       pgtype.Text{String: c.Contact.Email, Valid: true},
		Phone:       pgtype.Text{String: c.Contact.Phone, Valid: true},
		ContactName: pgtype.Text{String: c.Contact.ResposableName, Valid: true},

		CnpjCpf:    pgtype.Text{String: c.CnpjOrCpf, Valid: true},
		ClientType: pgstore.ClientType(c.ClientType),

		PostalCode:   pgtype.Text{String: c.Address.PostalCode, Valid: true},
		Neighborhood: pgtype.Text{String: c.Address.Neighborhood, Valid: true},
		Country:      pgtype.Text{String: c.Address.Country, Valid: true},
		State:        pgtype.Text{String: c.Address.State, Valid: true},
		City:         pgtype.Text{String: c.Address.City, Valid: true},
		Street:       pgtype.Text{String: c.Address.Street, Valid: true},
		Number:       pgtype.Text{String: c.Address.Number, Valid: true},
		Complement:   pgtype.Text{String: c.Address.Complement, Valid: true},

		Latitude:  pgtype.Float8{Float64: c.Address.Latitude, Valid: true},
		Longitude: pgtype.Float8{Float64: c.Address.Longitude, Valid: true},
	}

	if err := u.db.UpdateClientQuery(ctx, arg); err != nil {
		return err
	}

	return nil
}
func (u *postgresClientsRepository) DeleteClient(id uuid.UUID, ctx context.Context) error {
	if err := u.db.DeleteClientQuery(ctx, id); err != nil {
		return err
	}

	return nil
}
