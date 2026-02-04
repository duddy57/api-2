package repository

import (
	"context"
	"olidesk-api-2/internal/domains"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(*domains.User, context.Context) (uuid.UUID, error)
	FindByID(uuid.UUID, context.Context) (*domains.User, error)
	FindByEmail(string, context.Context) (*domains.User, error)
	Update(*domains.User, context.Context) error
	Delete(uuid.UUID, context.Context) error
	GetMembers(context.Context) ([]*domains.Member, error)
}

type ClientRepository interface {
	SaveClient(*domains.Client, context.Context) (uuid.UUID, error)
	FindClientByID(uuid.UUID, context.Context) (*domains.Client, error)
	ListClients(context.Context) ([]*domains.Client, error)
	UpdateClient(*domains.Client, context.Context) error
	DeleteClient(uuid.UUID, context.Context) error
}

type FormRepository interface {
	SaveForm(*domains.Atendimentos, context.Context) (uuid.UUID, error)
	FindFormByID(uuid.UUID, context.Context) (*domains.Atendimentos, error)
	ListForms(context.Context) ([]*domains.Atendimentos, error)
	UpdateForm(*domains.Atendimentos, context.Context) error
	DeleteForm(uuid.UUID, context.Context) error
}
