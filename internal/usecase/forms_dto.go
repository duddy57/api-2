package usecase

import (
	"time"

	"github.com/google/uuid"
)

type CreateFormInput struct {
	TecnicoResponsavelId []uuid.UUID `json:"tecnico_responsavel"`
	DataDeAbertura       time.Time   `json:"data_de_abertura"`
	ClienteId            uuid.UUID   `json:"cliente_id"`
	SolicitedBy          string      `json:"solicited_by"`
	DifficultyLevel      string      `json:"difficulty_level"`
	DefectDescription    string      `json:"defect_description"`
	SolutionDescription  string      `json:"solution_description"`
}

type UpdateFormInput struct {
	ID                   uuid.UUID   `json:"id"`
	DataDeAbertura       time.Time   `json:"data_de_abertura"`
	TecnicoResponsavelId []uuid.UUID `json:"tecnico_responsavel"`
	ClienteId            uuid.UUID   `json:"cliente_id"`
	SolicitedBy          string      `json:"solicited_by"`
	DifficultyLevel      string      `json:"difficulty_level"`
	DefectDescription    string      `json:"defect_description"`
	SolutionDescription  string      `json:"solution_description"`
}

type ListFormsOutput struct {
	Forms []FormsOutput `json:"forms"`
}

type GetFormsOutput struct {
	Form FormsOutput `json:"form"`
}

type FormsOutput struct {
	ID uuid.UUID `json:"id"`

	DataDeAbertura       time.Time  `json:"data_de_abertura"`
	TecnicoResponsavelId []Tecnicos `json:"tecnicos_responsaveis"`
	ClienteId            Client     `json:"cliente_id"`
	SolicitedBy          string     `json:"solicited_by"`
	DifficultyLevel      string     `json:"difficulty_level"`
	DefectDescription    string     `json:"defect_description"`
	SolutionDescription  string     `json:"solution_description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tecnicos struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Client struct {
	ID         uuid.UUID `json:"id"`
	ClientName string    `json:"client_name"`
}
