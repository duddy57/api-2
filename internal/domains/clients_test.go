package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestClient_Validate tests the Validate method with various scenarios
func TestClient_Validate(t *testing.T) {
	tests := []struct {
		name    string
		client  Client
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid client - all fields correct",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: false,
		},
		{
			name: "valid client - fisica type",
			client: Client{
				Address: Address{
					City:       "Rio de Janeiro",
					Complement: "Apto 501",
					Country:    "Brasil",
					Latitude:   -22.906847,
					Longitude:  -43.172897,
					Number:     "100",
					PostalCode: "20040-020",
					State:      "RJ",
					Street:     "Av. Rio Branco",
				},
				ClientName: "João da Silva",
				ClientType: "fisica",
				CnpjOrCpf:  "123.456.789-00",
				Contact: ContactPerson{
					Email:          "joao.silva@email.com",
					Phone:          "+55 21 98765-4321",
					ResposableName: "João da Silva",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid client - empty client name",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "client name",
		},
		{
			name: "invalid client - empty client type",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "client type",
		},
		{
			name: "invalid client - invalid client type",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "invalido",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "client type",
		},
		{
			name: "invalid client - empty cnpj or cpf",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "cnpj or cpf",
		},
		{
			name: "invalid client - empty contact name",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "",
				},
			},
			wantErr: true,
			errMsg:  "contact name",
		},
		{
			name: "invalid client - empty contact email",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "contact email",
		},
		{
			name: "invalid client - invalid contact email format",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "notanemail",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "contact email",
		},
		{
			name: "invalid client - empty contact phone",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "contact phone",
		},
		{
			name: "invalid client - empty postal code",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "postal code",
		},
		{
			name: "invalid client - empty country",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "country",
		},
		{
			name: "invalid client - empty state",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "state",
		},
		{
			name: "invalid client - empty city",
			client: Client{
				Address: Address{
					City:       "",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "city",
		},
		{
			name: "invalid client - empty street",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "1578",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "street",
		},
		{
			name: "invalid client - empty number",
			client: Client{
				Address: Address{
					City:       "São Paulo",
					Complement: "Sala 1203",
					Country:    "Brasil",
					Latitude:   -23.550520,
					Longitude:  -46.633308,
					Number:     "",
					PostalCode: "01310-200",
					State:      "SP",
					Street:     "Avenida Paulista",
				},
				ClientName: "Empresa Exemplo Tecnologia LTDA",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					Email:          "contato@empresaexemplo.com.br",
					Phone:          "+55 11 91234-5678",
					ResposableName: "Carlos Henrique Almeida",
				},
			},
			wantErr: true,
			errMsg:  "number",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.client.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestClient_Initialization tests entity field initialization
func TestClient_Initialization(t *testing.T) {
	client := &Client{
		ClientName: "Empresa Teste LTDA",
		ClientType: "juridica",
		CnpjOrCpf:  "12.345.678/0001-90",
		Contact: ContactPerson{
			ResposableName: "João Silva",
			Phone:          "+55 11 91234-5678",
			Email:          "joao@empresa.com",
		},
		Address: Address{
			PostalCode: "01310-200",
			Country:    "Brasil",
			State:      "SP",
			City:       "São Paulo",
			Street:     "Av. Paulista",
			Number:     "1000",
			Complement: "Sala 100",
			Latitude:   -23.550520,
			Longitude:  -46.633308,
		},
	}

	assert.Equal(t, "Empresa Teste LTDA", client.ClientName, "ClientName should be set correctly")
	assert.Equal(t, "juridica", client.ClientType, "ClientType should be set correctly")
	assert.Equal(t, "12.345.678/0001-90", client.CnpjOrCpf, "CnpjOrCpf should be set correctly")
	assert.Equal(t, "João Silva", client.Contact.ResposableName, "Contact name should be set correctly")
	assert.Equal(t, "joao@empresa.com", client.Contact.Email, "Contact email should be set correctly")
	assert.Equal(t, "+55 11 91234-5678", client.Contact.Phone, "Contact phone should be set correctly")
	assert.Equal(t, "01310-200", client.Address.PostalCode, "Postal code should be set correctly")
	assert.Equal(t, "São Paulo", client.Address.City, "City should be set correctly")
}

// TestClient_ClientType_EdgeCases tests edge cases for ClientType field
func TestClient_ClientType_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid juridica", value: "juridica", wantErr: false},
		{name: "valid fisica", value: "fisica", wantErr: false},
		{name: "empty string", value: "", wantErr: true},
		{name: "invalid type", value: "empresarial", wantErr: true},
		{name: "invalid type uppercase", value: "JURIDICA", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				ClientName: "Empresa Teste",
				ClientType: tt.value,
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					ResposableName: "João Silva",
					Phone:          "+55 11 91234-5678",
					Email:          "joao@empresa.com",
				},
				Address: Address{
					PostalCode: "01310-200",
					Country:    "Brasil",
					State:      "SP",
					City:       "São Paulo",
					Street:     "Av. Paulista",
					Number:     "1000",
				},
			}

			err := client.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestClient_ContactEmail_EdgeCases tests edge cases for Contact Email field
func TestClient_ContactEmail_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid email", value: "test@example.com", wantErr: false},
		{name: "valid email with subdomain", value: "user@mail.example.com", wantErr: false},
		{name: "empty string", value: "", wantErr: true},
		{name: "invalid email format", value: "notanemail", wantErr: true},
		{name: "email without @", value: "testexample.com", wantErr: true},
		{name: "email without domain", value: "test@", wantErr: true},
		{name: "email with spaces", value: "test @example.com", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				ClientName: "Empresa Teste",
				ClientType: "juridica",
				CnpjOrCpf:  "12.345.678/0001-90",
				Contact: ContactPerson{
					ResposableName: "João Silva",
					Phone:          "+55 11 91234-5678",
					Email:          tt.value,
				},
				Address: Address{
					PostalCode: "01310-200",
					Country:    "Brasil",
					State:      "SP",
					City:       "São Paulo",
					Street:     "Av. Paulista",
					Number:     "1000",
				},
			}

			err := client.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestClient_Address_EdgeCases tests edge cases for Address fields
func TestClient_Address_EdgeCases(t *testing.T) {
	baseClient := Client{
		ClientName: "Empresa Teste",
		ClientType: "juridica",
		CnpjOrCpf:  "12.345.678/0001-90",
		Contact: ContactPerson{
			ResposableName: "João Silva",
			Phone:          "+55 11 91234-5678",
			Email:          "joao@empresa.com",
		},
		Address: Address{
			PostalCode: "01310-200",
			Country:    "Brasil",
			State:      "SP",
			City:       "São Paulo",
			Street:     "Av. Paulista",
			Number:     "1000",
			Complement: "Sala 100",
		},
	}

	tests := []struct {
		name     string
		modifier func(*Client)
		wantErr  bool
	}{
		{
			name:     "valid with all fields",
			modifier: func(c *Client) {},
			wantErr:  false,
		},
		{
			name: "valid without complement",
			modifier: func(c *Client) {
				c.Address.Complement = ""
			},
			wantErr: false,
		},
		{
			name: "invalid - no postal code",
			modifier: func(c *Client) {
				c.Address.PostalCode = ""
			},
			wantErr: true,
		},
		{
			name: "invalid - no country",
			modifier: func(c *Client) {
				c.Address.Country = ""
			},
			wantErr: true,
		},
		{
			name: "invalid - no state",
			modifier: func(c *Client) {
				c.Address.State = ""
			},
			wantErr: true,
		},
		{
			name: "invalid - no city",
			modifier: func(c *Client) {
				c.Address.City = ""
			},
			wantErr: true,
		},
		{
			name: "invalid - no street",
			modifier: func(c *Client) {
				c.Address.Street = ""
			},
			wantErr: true,
		},
		{
			name: "invalid - no number",
			modifier: func(c *Client) {
				c.Address.Number = ""
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of base client
			client := baseClient
			tt.modifier(&client)

			err := client.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
