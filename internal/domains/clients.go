package domains

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID         uuid.UUID     `json:"id"`
	ClientName string        `json:"client_name"`
	Contact    ContactPerson `json:"contact_person"`
	CnpjOrCpf  string        `json:"cnpj_or_cpf"`
	ClientType string        `json:"client_type"`
	Address    Address       `json:"address"`
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

func (c *Client) Validate() error {
	if c.ClientName == "" {
		return ErrInvalidClientName
	}
	if c.ClientType == "" {
		return ErrInvalidClientType
	}
	if c.ClientType != "fisica" && c.ClientType != "juridica" {
		return ErrInvalidClientType
	}
	if c.CnpjOrCpf == "" {
		return ErrInvalidCnpjOrCpf
	}
	// Validate Contact Person
	if c.Contact.ResposableName == "" {
		return ErrInvalidContactName
	}
	if c.Contact.Email == "" {
		return ErrInvalidContactEmail
	}
	if !emailRegex.MatchString(c.Contact.Email) {
		return ErrInvalidContactEmail
	}
	if c.Contact.Phone == "" {
		return ErrInvalidContactPhone
	}
	// Validate Address
	if c.Address.PostalCode == "" {
		return ErrInvalidPostalCode
	}
	if c.Address.Country == "" {
		return ErrInvalidCountry
	}
	if c.Address.State == "" {
		return ErrInvalidState
	}
	if c.Address.City == "" {
		return ErrInvalidCity
	}
	if c.Address.Street == "" {
		return ErrInvalidStreet
	}
	if c.Address.Number == "" {
		return ErrInvalidNumber
	}
	return nil
}
