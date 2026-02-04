package usecase

import (
	"time"

	"github.com/google/uuid"
)

type CreateClientInput struct {
	ClientName string        `json:"client_name"`
	Contact    ContactPerson `json:"contact_person"`
	CnpjOrCpf  string        `json:"cnpj_or_cpf"`
	Address    Address       `json:"address"`
	ClientType string        `json:"client_type"`
}
type UpdateClientInput struct {
	ClientName string        `json:"client_name"`
	Contact    ContactPerson `json:"contact_person"`
	CnpjOrCpf  string        `json:"cnpj_or_cpf"`
	Address    Address       `json:"address"`
	ClientType string        `json:"client_type"`
}

type GetClientOutput struct {
	Client *ClientOutput `json:"clients"`
}

type ListClientOutput struct {
	Clients []*ClientOutput `json:"clients"`
}

type ClientOutput struct {
	ID         uuid.UUID     `json:"id"`
	ClientName string        `json:"client_name"`
	Contact    ContactPerson `json:"contact_person"`
	CnpjOrCpf  string        `json:"cnpj_or_cpf"`
	Address    Address       `json:"address"`
	ClientType string        `json:"client_type"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

type Address struct {
	PostalCode   string  `json:"postal_code"`
	Neighborhood string  `json:"neighborhood"`
	Country      string  `json:"country"`
	State        string  `json:"state"`
	City         string  `json:"city"`
	Street       string  `json:"street"`
	Number       string  `json:"number"`
	Complement   string  `json:"complement"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
type ContactPerson struct {
	ResposableName string `json:"responsible_name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}
