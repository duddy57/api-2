-- +goose Up
-- +goose StatementBegin
CREATE TYPE difficulty_level AS ENUM ('low', 'medium', 'high');

CREATE TABLE IF NOT EXISTS forms (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    client_id UUID NOT NULL REFERENCES clients(id),
    solicited_name VARCHAR(50) NOT NULL,
    difficulty_level difficulty_level NOT NULL DEFAULT 'low',
    defect_description VARCHAR(255),
    solution_description VARCHAR(255),
    occurred_at TIMESTAMPTZ NOT NULL,

    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS form_tecnico (
    id UUID PRIMARY KEY DEFAULT uuidv7(),

    member_id UUID NOT NULL REFERENCES members(id) ON DELETE SET NULL,
    form_id UUID NOT NULL REFERENCES forms(id) ON DELETE CASCADE,

    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS form_tecnico;
DROP TABLE IF EXISTS forms;
-- +goose StatementEnd
