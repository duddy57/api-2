package domains

import (
	"time"

	"github.com/google/uuid"
)

type Atendimentos struct {
	ID uuid.UUID `json:"id"`

	DataDeAbertura       time.Time  `json:"data_de_abertura"`
	TecnicoResponsavelId []Member   `json:"tecnicos_responsavel"`
	Cliente              ClientForm `json:"cliente"`
	SolicitedBy          string     `json:"solicited_by"`
	DifficultyLevel      string     `json:"difficulty_level"`
	DefectDescription    string     `json:"defect_description"`
	SolutionDescription  string     `json:"solution_description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClientForm struct {
	ID         uuid.UUID `json:"id"`
	ClientName string    `json:"client_name"`
}

func (u *Atendimentos) Validate() error {
	if u.DefectDescription == "" {
		return ErrInvalidDefectDescription
	}
	if u.DifficultyLevel == "" {
		return ErrInvalidDifficultyLevel
	}
	if u.SolicitedBy == "" {
		return ErrInvalidSolicitedBy
	}
	if u.Cliente.ID == uuid.Nil {
		return ErrInvalidClienteId
	}
	if len(u.TecnicoResponsavelId) == 0 {
		return ErrInvalidTecnicoResponsavelId
	}
	if u.DataDeAbertura.IsZero() {
		return ErrInvalidDataDeAbertura
	}

	if u.SolutionDescription == "" {
		return ErrInvalidSolutionDescription
	}

	return nil
}
