package usecase

import (
	"context"
	"olidesk-api-2/internal/domains"
	"olidesk-api-2/internal/repository"
	"olidesk-api-2/internal/utils/location"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ClientUseCase interface {
	CreateClient(CreateClientInput, context.Context) (uuid.UUID, error)
	GetClient(uuid.UUID, context.Context) (*ClientOutput, error)
	ListClient(context.Context) ([]*ClientOutput, error)
	UpdateClient(uuid.UUID, UpdateClientInput, context.Context) error
	DeleteClient(uuid.UUID, context.Context) error
}

type clientService struct {
	repo repository.ClientRepository
	l    *zap.Logger
}

func NewClientService(repo repository.ClientRepository, l *zap.Logger) ClientUseCase {
	return &clientService{repo: repo, l: l}
}

func (c *clientService) CreateClient(p CreateClientInput, ctx context.Context) (uuid.UUID, error) {
	lat, lng, err := location.GeocodeAddress(
		ctx,
		p.Address.Street,
		p.Address.Number,
		p.Address.Neighborhood,
		p.Address.City,
		p.Address.State,
		p.Address.PostalCode,
		p.Address.Country,
	)
	if err != nil {
		c.l.Error("error geocoding address", zap.Error(err))
		return uuid.Nil, err
	}
	p.Address.Latitude = lat
	p.Address.Longitude = lng

	id, err := c.repo.SaveClient(&domains.Client{
		ClientName: p.ClientName,
		CnpjOrCpf:  p.CnpjOrCpf,
		ClientType: p.ClientType,
		Contact: domains.ContactPerson{
			Email:          p.Contact.Email,
			Phone:          p.Contact.Phone,
			ResposableName: p.Contact.ResposableName,
		},
		Address: domains.Address{
			PostalCode:   p.Address.PostalCode,
			Neighborhood: p.Address.Neighborhood,
			Country:      p.Address.Country,
			State:        p.Address.State,
			City:         p.Address.City,
			Street:       p.Address.Street,
			Number:       p.Address.Number,
			Complement:   p.Address.Complement,
			Latitude:     lat,
			Longitude:    lng,
		},
	}, ctx)

	if err != nil {
		c.l.Error("error saving client", zap.Error(err))
		return uuid.Nil, err
	}
	return id, nil
}
func (c *clientService) DeleteClient(id uuid.UUID, ctx context.Context) error {
	if err := c.repo.DeleteClient(id, ctx); err != nil {
		c.l.Error("error deleting client", zap.Error(err))
		return err
	}
	return nil
}
func (c *clientService) GetClient(id uuid.UUID, ctx context.Context) (*ClientOutput, error) {
	client, err := c.repo.FindClientByID(id, ctx)
	if err != nil {
		c.l.Error("error getting client", zap.Error(err))
		return nil, err
	}
	return &ClientOutput{
		ClientName: client.ClientName,
		CnpjOrCpf:  client.CnpjOrCpf,
		ClientType: client.ClientType,
		Contact: ContactPerson{
			Email:          client.Contact.Email,
			Phone:          client.Contact.Phone,
			ResposableName: client.Contact.ResposableName,
		},
		Address: Address{
			PostalCode:   client.Address.PostalCode,
			Neighborhood: client.Address.Neighborhood,
			Country:      client.Address.Country,
			State:        client.Address.State,
			City:         client.Address.City,
			Street:       client.Address.Street,
			Number:       client.Address.Number,
			Complement:   client.Address.Complement,
			Latitude:     client.Address.Latitude,
			Longitude:    client.Address.Longitude,
		},
	}, nil
}
func (c *clientService) ListClient(ctx context.Context) ([]*ClientOutput, error) {
	clientData, err := c.repo.ListClients(ctx)
	if err != nil {
		c.l.Error("error listing clients", zap.Error(err))
		return nil, err
	}

	clientList := make([]*ClientOutput, 0, len(clientData))
	for _, cl := range clientData {
		clientList = append(clientList, &ClientOutput{
			ID:         cl.ID,
			ClientName: cl.ClientName,
			CnpjOrCpf:  cl.CnpjOrCpf,
			ClientType: cl.ClientType,
			Contact: ContactPerson{
				Email:          cl.Contact.Email,
				Phone:          cl.Contact.Phone,
				ResposableName: cl.Contact.ResposableName,
			},
			Address: Address{
				PostalCode:   cl.Address.PostalCode,
				Neighborhood: cl.Address.Neighborhood,
				Country:      cl.Address.Country,
				State:        cl.Address.State,
				City:         cl.Address.City,
				Street:       cl.Address.Street,
				Number:       cl.Address.Number,
				Complement:   cl.Address.Complement,
				Latitude:     cl.Address.Latitude,
				Longitude:    cl.Address.Longitude,
			},
			CreatedAt: cl.CreatedAt.UTC(),
			UpdatedAt: cl.UpdatedAt.UTC(),
		})
	}

	return clientList, nil
}
func (c *clientService) UpdateClient(id uuid.UUID, cl UpdateClientInput, ctx context.Context) error {
	client, err := c.repo.FindClientByID(id, ctx)
	if err != nil {
		return err
	}

	if cl.ClientName != "" {
		client.ClientName = cl.ClientName
	}
	if cl.CnpjOrCpf != "" {
		client.CnpjOrCpf = cl.CnpjOrCpf
	}
	if cl.ClientType != "" {
		client.ClientType = cl.ClientType
	}

	if cl.Contact.Email != "" {
		client.Contact.Email = cl.Contact.Email
	}
	if cl.Contact.Phone != "" {
		client.Contact.Phone = cl.Contact.Phone
	}
	if cl.Contact.ResposableName != "" {
		client.Contact.ResposableName = cl.Contact.ResposableName
	}

	if cl.Address.Neighborhood != "" {
		client.Address.Neighborhood = cl.Address.Neighborhood
	}
	if cl.Address.PostalCode != "" {
		client.Address.PostalCode = cl.Address.PostalCode
	}
	if cl.Address.Country != "" {
		client.Address.Country = cl.Address.Country
	}
	if cl.Address.State != "" {
		client.Address.State = cl.Address.State
	}
	if cl.Address.City != "" {
		client.Address.City = cl.Address.City
	}
	if cl.Address.Street != "" {
		client.Address.Street = cl.Address.Street
	}
	if cl.Address.Number != "" {
		client.Address.Number = cl.Address.Number
	}
	if cl.Address.Complement != "" {
		client.Address.Complement = cl.Address.Complement
	}

	if cl.Address.Street != "" &&
		cl.Address.Number != "" &&
		cl.Address.City != "" &&
		cl.Address.State != "" &&
		cl.Address.PostalCode != "" &&
		cl.Address.Country != "" &&
		cl.Address.Neighborhood != "" {
		lat, lng, err := location.GeocodeAddress(
			ctx,
			cl.Address.Street,
			cl.Address.Neighborhood,
			cl.Address.Number,
			cl.Address.City,
			cl.Address.State,
			cl.Address.PostalCode,
			cl.Address.Country,
		)
		if err != nil {
			c.l.Error("error geocoding address", zap.Error(err))
			return err
		}
		client.Address.Latitude = lat
		client.Address.Longitude = lng
	}

	if err := c.repo.UpdateClient(client, ctx); err != nil {
		c.l.Error("error updating client", zap.Error(err))
		return err
	}
	return nil
}
