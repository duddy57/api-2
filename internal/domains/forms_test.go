package domains

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestAtendimentos_Validate tests the Validate method with various scenarios
func TestAtendimentos_Validate(t *testing.T) {
	validMember := []Member{
		{
			ID:    uuid.New(),
			Name:  "Carlos Silva",
			Email: "carlos@example.com",
			Role:  "tecnico",
		},
	}

	validClient := ClientForm{
		ID:         uuid.New(),
		ClientName: "Empresa Exemplo LTDA",
	}

	validDate := time.Now()

	tests := []struct {
		name        string
		atendimento Atendimentos
		wantErr     bool
		expectedErr error
	}{
		{
			name: "valid atendimento - all fields correct",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "João da Silva",
				DifficultyLevel:      "medio",
				DefectDescription:    "Sistema apresentando lentidão ao processar requisições",
				SolutionDescription:  "Otimizado banco de dados e aumentado cache",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr: false,
		},
		{
			name: "valid atendimento - difficulty level baixo",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "Maria Santos",
				DifficultyLevel:      "baixo",
				DefectDescription:    "Erro de autenticação no login",
				SolutionDescription:  "Resetado as credenciais do usuário",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr: false,
		},
		{
			name: "valid atendimento - difficulty level alto",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "Pedro Oliveira",
				DifficultyLevel:      "alto",
				DefectDescription:    "Sistema completamente fora do ar",
				SolutionDescription:  "Recuperado servidor de backup e restaurado dados",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr: false,
		},
		{
			name: "valid atendimento - multiple technicians",
			atendimento: Atendimentos{
				ID:             uuid.New(),
				DataDeAbertura: validDate,
				TecnicoResponsavelId: []Member{
					{
						ID:    uuid.New(),
						Name:  "Carlos Silva",
						Email: "carlos@example.com",
						Role:  "tecnico",
					},
					{
						ID:    uuid.New(),
						Name:  "Ana Paula",
						Email: "ana@example.com",
						Role:  "tecnico",
					},
				},
				Cliente:             validClient,
				SolicitedBy:         "Gerente TI",
				DifficultyLevel:     "alto",
				DefectDescription:   "Integração com sistema externo falhando",
				SolutionDescription: "Reconfigurado API e atualizado certificados",
				CreatedAt:           validDate,
				UpdatedAt:           validDate,
			},
			wantErr: false,
		},
		{
			name: "invalid atendimento - empty defect description",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "João da Silva",
				DifficultyLevel:      "medio",
				DefectDescription:    "",
				SolutionDescription:  "Solução aplicada",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr:     true,
			expectedErr: ErrInvalidDefectDescription,
		},
		{
			name: "invalid atendimento - empty difficulty level",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "João da Silva",
				DifficultyLevel:      "",
				DefectDescription:    "Problema identificado",
				SolutionDescription:  "Solução aplicada",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr:     true,
			expectedErr: ErrInvalidDifficultyLevel,
		},
		{
			name: "invalid atendimento - empty solicited by",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "",
				DifficultyLevel:      "medio",
				DefectDescription:    "Problema identificado",
				SolutionDescription:  "Solução aplicada",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr:     true,
			expectedErr: ErrInvalidSolicitedBy,
		},
		{
			name: "invalid atendimento - nil client ID",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente: ClientForm{
					ID:         uuid.Nil,
					ClientName: "Empresa Exemplo",
				},
				SolicitedBy:         "João da Silva",
				DifficultyLevel:     "medio",
				DefectDescription:   "Problema identificado",
				SolutionDescription: "Solução aplicada",
				CreatedAt:           validDate,
				UpdatedAt:           validDate,
			},
			wantErr:     true,
			expectedErr: ErrInvalidClienteId,
		},
		{
			name: "invalid atendimento - empty technician list",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: []Member{},
				Cliente:              validClient,
				SolicitedBy:          "João da Silva",
				DifficultyLevel:      "medio",
				DefectDescription:    "Problema identificado",
				SolutionDescription:  "Solução aplicada",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr:     true,
			expectedErr: ErrInvalidTecnicoResponsavelId,
		},
		{
			name: "invalid atendimento - zero opening date",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       time.Time{},
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "João da Silva",
				DifficultyLevel:      "medio",
				DefectDescription:    "Problema identificado",
				SolutionDescription:  "Solução aplicada",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr:     true,
			expectedErr: ErrInvalidDataDeAbertura,
		},
		{
			name: "invalid atendimento - empty solution description",
			atendimento: Atendimentos{
				ID:                   uuid.New(),
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "João da Silva",
				DifficultyLevel:      "medio",
				DefectDescription:    "Problema identificado",
				SolutionDescription:  "",
				CreatedAt:            validDate,
				UpdatedAt:            validDate,
			},
			wantErr:     true,
			expectedErr: ErrInvalidSolutionDescription,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.atendimento.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					assert.Equal(t, tt.expectedErr, err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestAtendimentos_Initialization tests entity field initialization
func TestAtendimentos_Initialization(t *testing.T) {
	id := uuid.New()
	clientID := uuid.New()
	technicianID := uuid.New()
	openDate := time.Now()
	createdAt := time.Now()
	updatedAt := time.Now()

	atendimento := &Atendimentos{
		ID:             id,
		DataDeAbertura: openDate,
		TecnicoResponsavelId: []Member{
			{
				ID:    technicianID,
				Name:  "Carlos Silva",
				Email: "carlos@example.com",
				Role:  "tecnico",
			},
		},
		Cliente: ClientForm{
			ID:         clientID,
			ClientName: "Empresa Exemplo LTDA",
		},
		SolicitedBy:         "João da Silva",
		DifficultyLevel:     "medio",
		DefectDescription:   "Sistema apresentando lentidão",
		SolutionDescription: "Otimizado banco de dados",
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
	}

	assert.Equal(t, id, atendimento.ID, "ID should be set correctly")
	assert.Equal(t, openDate, atendimento.DataDeAbertura, "DataDeAbertura should be set correctly")
	assert.Len(t, atendimento.TecnicoResponsavelId, 1, "Should have 1 technician")
	assert.Equal(t, technicianID, atendimento.TecnicoResponsavelId[0].ID, "Technician ID should be set correctly")
	assert.Equal(t, "Carlos Silva", atendimento.TecnicoResponsavelId[0].Name, "Technician name should be set correctly")
	assert.Equal(t, clientID, atendimento.Cliente.ID, "Client ID should be set correctly")
	assert.Equal(t, "Empresa Exemplo LTDA", atendimento.Cliente.ClientName, "Client name should be set correctly")
	assert.Equal(t, "João da Silva", atendimento.SolicitedBy, "SolicitedBy should be set correctly")
	assert.Equal(t, "medio", atendimento.DifficultyLevel, "DifficultyLevel should be set correctly")
	assert.Equal(t, "Sistema apresentando lentidão", atendimento.DefectDescription, "DefectDescription should be set correctly")
	assert.Equal(t, "Otimizado banco de dados", atendimento.SolutionDescription, "SolutionDescription should be set correctly")
	assert.Equal(t, createdAt, atendimento.CreatedAt, "CreatedAt should be set correctly")
	assert.Equal(t, updatedAt, atendimento.UpdatedAt, "UpdatedAt should be set correctly")
}

// TestAtendimentos_DefectDescription_EdgeCases tests edge cases for DefectDescription field
func TestAtendimentos_DefectDescription_EdgeCases(t *testing.T) {
	validMember := []Member{{ID: uuid.New(), Name: "Tech", Email: "tech@example.com", Role: "tecnico"}}
	validClient := ClientForm{ID: uuid.New(), ClientName: "Cliente"}
	validDate := time.Now()

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid description", value: "Sistema apresentando erro 500", wantErr: false},
		{name: "empty string", value: "", wantErr: true},
		{name: "long description", value: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.", wantErr: false},
		{name: "single character", value: "X", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atendimento := &Atendimentos{
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "Solicitante",
				DifficultyLevel:      "medio",
				DefectDescription:    tt.value,
				SolutionDescription:  "Solução aplicada",
			}

			err := atendimento.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestAtendimentos_DifficultyLevel_EdgeCases tests edge cases for DifficultyLevel field
func TestAtendimentos_DifficultyLevel_EdgeCases(t *testing.T) {
	validMember := []Member{{ID: uuid.New(), Name: "Tech", Email: "tech@example.com", Role: "tecnico"}}
	validClient := ClientForm{ID: uuid.New(), ClientName: "Cliente"}
	validDate := time.Now()

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid level - baixo", value: "baixo", wantErr: false},
		{name: "valid level - medio", value: "medio", wantErr: false},
		{name: "valid level - alto", value: "alto", wantErr: false},
		{name: "empty string", value: "", wantErr: true},
		{name: "invalid level", value: "extremo", wantErr: false}, // No validation for specific values
		{name: "uppercase", value: "ALTO", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atendimento := &Atendimentos{
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "Solicitante",
				DifficultyLevel:      tt.value,
				DefectDescription:    "Descrição do defeito",
				SolutionDescription:  "Solução aplicada",
			}

			err := atendimento.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestAtendimentos_SolutionDescription_EdgeCases tests edge cases for SolutionDescription field
func TestAtendimentos_SolutionDescription_EdgeCases(t *testing.T) {
	validMember := []Member{{ID: uuid.New(), Name: "Tech", Email: "tech@example.com", Role: "tecnico"}}
	validClient := ClientForm{ID: uuid.New(), ClientName: "Cliente"}
	validDate := time.Now()

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid solution", value: "Problema resolvido com atualização", wantErr: false},
		{name: "empty string", value: "", wantErr: true},
		{name: "long solution", value: "Solução detalhada incluindo múltiplas etapas de correção, testes realizados e validação final com o cliente para garantir que o problema foi completamente resolvido.", wantErr: false},
		{name: "single character", value: "X", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atendimento := &Atendimentos{
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              validClient,
				SolicitedBy:          "Solicitante",
				DifficultyLevel:      "medio",
				DefectDescription:    "Descrição do defeito",
				SolutionDescription:  tt.value,
			}

			err := atendimento.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestAtendimentos_TecnicoResponsavelId_EdgeCases tests edge cases for TecnicoResponsavelId field
func TestAtendimentos_TecnicoResponsavelId_EdgeCases(t *testing.T) {
	validClient := ClientForm{ID: uuid.New(), ClientName: "Cliente"}
	validDate := time.Now()

	tests := []struct {
		name    string
		value   []Member
		wantErr bool
	}{
		{
			name: "valid - single technician",
			value: []Member{
				{ID: uuid.New(), Name: "Carlos", Email: "carlos@example.com", Role: "tecnico"},
			},
			wantErr: false,
		},
		{
			name: "valid - multiple technicians",
			value: []Member{
				{ID: uuid.New(), Name: "Carlos", Email: "carlos@example.com", Role: "tecnico"},
				{ID: uuid.New(), Name: "Ana", Email: "ana@example.com", Role: "tecnico"},
				{ID: uuid.New(), Name: "Pedro", Email: "pedro@example.com", Role: "senior"},
			},
			wantErr: false,
		},
		{
			name:    "invalid - empty list",
			value:   []Member{},
			wantErr: true,
		},
		{
			name:    "invalid - nil list",
			value:   nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atendimento := &Atendimentos{
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: tt.value,
				Cliente:              validClient,
				SolicitedBy:          "Solicitante",
				DifficultyLevel:      "medio",
				DefectDescription:    "Descrição do defeito",
				SolutionDescription:  "Solução aplicada",
			}

			err := atendimento.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestClientForm_EdgeCases tests edge cases for ClientForm struct
func TestClientForm_EdgeCases(t *testing.T) {
	validMember := []Member{{ID: uuid.New(), Name: "Tech", Email: "tech@example.com", Role: "tecnico"}}
	validDate := time.Now()

	tests := []struct {
		name    string
		client  ClientForm
		wantErr bool
	}{
		{
			name: "valid client",
			client: ClientForm{
				ID:         uuid.New(),
				ClientName: "Empresa Exemplo LTDA",
			},
			wantErr: false,
		},
		{
			name: "invalid - nil UUID",
			client: ClientForm{
				ID:         uuid.Nil,
				ClientName: "Empresa Exemplo LTDA",
			},
			wantErr: true,
		},
		{
			name: "valid - empty client name (not validated in ClientForm)",
			client: ClientForm{
				ID:         uuid.New(),
				ClientName: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atendimento := &Atendimentos{
				DataDeAbertura:       validDate,
				TecnicoResponsavelId: validMember,
				Cliente:              tt.client,
				SolicitedBy:          "Solicitante",
				DifficultyLevel:      "medio",
				DefectDescription:    "Descrição do defeito",
				SolutionDescription:  "Solução aplicada",
			}

			err := atendimento.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
