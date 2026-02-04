package usecase

import (
	"context"
	"olidesk-api-2/internal/domains"
	"olidesk-api-2/internal/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type formService struct {
	repo repository.FormRepository
	l    *zap.Logger
}

type FormsUseCase interface {
	CreateForm(CreateFormInput, context.Context) (uuid.UUID, error)
	GetForm(uuid.UUID, context.Context) (*GetFormsOutput, error)
	UpdateForm(uuid.UUID, UpdateFormInput, context.Context) error
	DeleteForm(uuid.UUID, context.Context) error
	ListForms(context.Context) (*ListFormsOutput, error)
}

func NewFormService(repo repository.FormRepository, l *zap.Logger) FormsUseCase {
	return &formService{
		repo: repo,
		l:    l,
	}
}

func (f *formService) CreateForm(p CreateFormInput, ctx context.Context) (uuid.UUID, error) {
	tecnicos := make([]domains.Member, 0, len(p.TecnicoResponsavelId))
	for _, tecID := range p.TecnicoResponsavelId {
		tecnicos = append(tecnicos, domains.Member{
			ID: tecID,
		})
	}

	id, err := f.repo.SaveForm(&domains.Atendimentos{
		TecnicoResponsavelId: tecnicos,
		Cliente: domains.ClientForm{
			ID: p.ClienteId,
		},
		DataDeAbertura:      p.DataDeAbertura,
		SolicitedBy:         p.SolicitedBy,
		DifficultyLevel:     p.DifficultyLevel,
		DefectDescription:   p.DefectDescription,
		SolutionDescription: p.SolutionDescription,
	}, ctx)
	if err != nil {
		f.l.Error("error creating form", zap.Error(err))
		return uuid.Nil, err
	}
	return id, nil
}
func (f *formService) GetForm(id uuid.UUID, ctx context.Context) (*GetFormsOutput, error) {
	form, err := f.repo.FindFormByID(id, ctx)
	if err != nil {
		f.l.Error("error getting form", zap.Error(err))
		return nil, err
	}

	tecnicos := make([]Tecnicos, 0, len(form.TecnicoResponsavelId))
	for _, tec := range form.TecnicoResponsavelId {
		tecnicos = append(tecnicos, Tecnicos{
			ID:   tec.ID,
			Name: tec.Name,
		})
	}

	return &GetFormsOutput{
		Form: FormsOutput{
			ID:                   form.ID,
			TecnicoResponsavelId: tecnicos,
			ClienteId: Client{
				ID:         form.Cliente.ID,
				ClientName: form.Cliente.ClientName,
			},
			SolicitedBy:         form.SolicitedBy,
			DifficultyLevel:     form.DifficultyLevel,
			DefectDescription:   form.DefectDescription,
			SolutionDescription: form.SolutionDescription,
			DataDeAbertura:      form.DataDeAbertura,
			UpdatedAt:           form.UpdatedAt,
			CreatedAt:           form.CreatedAt,
		},
	}, nil
}
func (f *formService) UpdateForm(id uuid.UUID, input UpdateFormInput, ctx context.Context) error {
	form, err := f.repo.FindFormByID(id, ctx)
	if err != nil {
		f.l.Error("error getting form", zap.Error(err))
		return err
	}

	if len(input.TecnicoResponsavelId) > 0 {
		tecnicos := make([]domains.Member, 0, len(input.TecnicoResponsavelId))
		for _, tecID := range input.TecnicoResponsavelId {
			tecnicos = append(tecnicos, domains.Member{ID: tecID})
		}
		form.TecnicoResponsavelId = tecnicos
	}
	if input.ClienteId != uuid.Nil {
		form.Cliente.ID = input.ClienteId
	}
	if input.SolicitedBy != "" {
		form.SolicitedBy = input.SolicitedBy
	}
	if input.DifficultyLevel != "" {
		form.DifficultyLevel = input.DifficultyLevel
	}
	if input.DefectDescription != "" {
		form.DefectDescription = input.DefectDescription
	}
	if input.SolutionDescription != "" {
		form.SolutionDescription = input.SolutionDescription
	}
	if err := f.repo.UpdateForm(form, ctx); err != nil {
		f.l.Error("error updating form", zap.Error(err))
		return err
	}

	return nil
}
func (f *formService) DeleteForm(id uuid.UUID, ctx context.Context) error {
	if err := f.repo.DeleteForm(id, ctx); err != nil {
		f.l.Error("error deleting form", zap.Error(err))
		return err
	}

	return nil
}
func (f *formService) ListForms(ctx context.Context) (*ListFormsOutput, error) {
	formData, err := f.repo.ListForms(ctx)
	if err != nil {
		f.l.Error("error listing forms", zap.Error(err))
		return nil, err
	}

	formList := make([]FormsOutput, 0, len(formData))
	for _, fl := range formData {
		tecnicos := make([]Tecnicos, 0, len(fl.TecnicoResponsavelId))
		for _, tec := range fl.TecnicoResponsavelId {
			tecnicos = append(tecnicos, Tecnicos{
				ID:   tec.ID,
				Name: tec.Name,
			})
		}

		formList = append(formList, FormsOutput{
			ID:                   fl.ID,
			TecnicoResponsavelId: tecnicos,
			ClienteId: Client{
				ID:         fl.Cliente.ID,
				ClientName: fl.Cliente.ClientName,
			},
			SolicitedBy:         fl.SolicitedBy,
			DifficultyLevel:     fl.DifficultyLevel,
			DefectDescription:   fl.DefectDescription,
			SolutionDescription: fl.SolutionDescription,
			DataDeAbertura:      fl.DataDeAbertura,
			UpdatedAt:           fl.UpdatedAt.UTC(),
			CreatedAt:           fl.CreatedAt.UTC(),
		})
	}

	return &ListFormsOutput{Forms: formList}, nil
}
