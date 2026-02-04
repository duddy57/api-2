package repository

import (
	"context"
	"fmt"
	"olidesk-api-2/internal/domains"
	"olidesk-api-2/internal/store/pgstore"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresFormRepository struct {
	db   *pgstore.Queries
	pool *pgxpool.Pool
}

func NewPostgresFormRepository(db *pgxpool.Pool) FormRepository {
	return &postgresFormRepository{db: pgstore.New(db), pool: db}
}

func (p *postgresFormRepository) SaveForm(input *domains.Atendimentos, ctx context.Context) (uuid.UUID, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to begin tx for RegisterTeam: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	qtx := p.db.WithTx(tx)

	result, err := qtx.CreateFormQuery(ctx, pgstore.CreateFormQueryParams{
		ClientID:            input.Cliente.ID,
		SolicitedName:       input.SolicitedBy,
		OccurredAt:          input.DataDeAbertura.UTC(),
		DifficultyLevel:     pgstore.DifficultyLevel(input.DifficultyLevel),
		DefectDescription:   pgtype.Text{String: input.DefectDescription, Valid: true},
		SolutionDescription: pgtype.Text{String: input.SolutionDescription, Valid: true},
	})

	tecnicos := make([]pgstore.CreateFormTecnicoQueryParams, len(input.TecnicoResponsavelId))
	for i, item := range input.TecnicoResponsavelId {
		tecnicos[i] = pgstore.CreateFormTecnicoQueryParams{
			FormID:   result,
			MemberID: item.ID,
		}
	}

	if _, err := qtx.CreateFormTecnicoQuery(ctx, tecnicos); err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}

	return result, nil
}
func (p *postgresFormRepository) FindFormByID(id uuid.UUID, ctx context.Context) (*domains.Atendimentos, error) {
	formDetails, err := p.db.GetFormByIdQuery(ctx, id)
	if err != nil {
		return nil, err
	}

	tecnicosRaw, err := p.db.GetFormTecnicosByFormID(ctx, id)
	if err != nil {
		return nil, err
	}

	tecnicosList := make([]domains.Member, 0, len(tecnicosRaw))
	for _, i := range tecnicosRaw {
		tecnicosList = append(tecnicosList, domains.Member{
			ID:    i.ID,
			Name:  i.UserName,
			Email: i.UserEmail,
		})
	}

	return &domains.Atendimentos{
		ID:             formDetails.ID,
		DataDeAbertura: formDetails.OccurredAt.UTC(),
		Cliente: domains.ClientForm{
			ID:         formDetails.ClientID,
			ClientName: formDetails.ClientName,
		},
		SolicitedBy:          formDetails.SolicitedName,
		DifficultyLevel:      string(formDetails.DifficultyLevel),
		DefectDescription:    formDetails.DefectDescription.String,
		SolutionDescription:  formDetails.SolutionDescription.String,
		CreatedAt:            formDetails.CreatedAt.Time,
		UpdatedAt:            formDetails.UpdatedAt.Time,
		TecnicoResponsavelId: tecnicosList,
	}, nil
}
func (p *postgresFormRepository) ListForms(ctx context.Context) ([]*domains.Atendimentos, error) {
	formDetails, err := p.db.GetFormsQuery(ctx)
	if err != nil {
		return nil, err
	}

	forms := make([]*domains.Atendimentos, 0, len(formDetails))
	for _, i := range formDetails {
		tecnicosRaw, err := p.db.GetFormTecnicosByFormID(ctx, i.ID)
		if err != nil {
			return nil, err
		}

		tecnicosList := make([]domains.Member, 0, len(tecnicosRaw))
		for _, t := range tecnicosRaw {
			tecnicosList = append(tecnicosList, domains.Member{
				ID:    t.ID,
				Name:  t.UserName,
				Email: t.UserEmail,
			})
		}

		forms = append(forms, &domains.Atendimentos{
			ID:             i.ID,
			DataDeAbertura: i.OccurredAt.UTC(),
			Cliente: domains.ClientForm{
				ID:         i.ClientID,
				ClientName: i.ClientName,
			},
			SolicitedBy:          i.SolicitedName,
			DifficultyLevel:      string(i.DifficultyLevel),
			DefectDescription:    i.DefectDescription.String,
			SolutionDescription:  i.SolutionDescription.String,
			CreatedAt:            i.CreatedAt.Time,
			UpdatedAt:            i.UpdatedAt.Time,
			TecnicoResponsavelId: tecnicosList,
		})
	}

	return forms, nil
}
func (p *postgresFormRepository) UpdateForm(input *domains.Atendimentos, ctx context.Context) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("pgstore: failed to begin tx for RegisterTeam: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	qtx := p.db.WithTx(tx)

	if err := qtx.UpdateFormQuery(ctx, pgstore.UpdateFormQueryParams{
		ClientID:            input.Cliente.ID,
		SolicitedName:       input.SolicitedBy,
		DifficultyLevel:     pgstore.DifficultyLevel(input.DifficultyLevel),
		DefectDescription:   pgtype.Text{String: input.DefectDescription, Valid: true},
		SolutionDescription: pgtype.Text{String: input.SolutionDescription, Valid: true},
		ID:                  input.ID,
	}); err != nil {
		return err
	}

	tecnicos := make([]pgstore.CreateFormTecnicoQueryParams, len(input.TecnicoResponsavelId))
	for i, item := range input.TecnicoResponsavelId {
		tecnicos[i] = pgstore.CreateFormTecnicoQueryParams{
			FormID:   input.ID,
			MemberID: item.ID,
		}
	}

	if _, err := qtx.CreateFormTecnicoQuery(ctx, tecnicos); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
func (p *postgresFormRepository) DeleteForm(id uuid.UUID, ctx context.Context) error {
	if err := p.db.DeleteFormQuery(ctx, id); err != nil {
		return err
	}

	return nil
}
