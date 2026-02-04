-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Tabela: members
-- Descrição: Membros do sistema com funções operacionais (técnicos e admins)
-- Relacionamento: 1:1 com users
-- Versão: 2.0
-- ============================================================================

CREATE TYPE member_role AS ENUM('tecnico_interno', 'tecnico_externo', 'administrador');

CREATE TABLE IF NOT EXISTS members (
    id UUID PRIMARY KEY DEFAULT uuidv7(),

    user_id UUID NOT NULL,


    role member_role NOT NULL DEFAULT 'tecnico_interno',

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT members_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT members_user_id_unique UNIQUE (user_id)
);

CREATE INDEX IF NOT EXISTS idx_members_user_id ON members(user_id);
CREATE INDEX IF NOT EXISTS idx_members_role ON members(role);
CREATE INDEX IF NOT EXISTS idx_members_active ON members(is_active) WHERE is_active = TRUE;
CREATE INDEX IF NOT EXISTS idx_members_role_active ON members(role, is_active) WHERE is_active = TRUE;
CREATE INDEX IF NOT EXISTS idx_members_created_at ON members(created_at DESC);


COMMENT ON TABLE members IS 'Membros operacionais do sistema (técnicos, auxiliares, estagiários e admins)';
COMMENT ON COLUMN members.id IS 'Identificador único do membro (UUID)';
COMMENT ON COLUMN members.user_id IS 'Referência ao usuário (relação 1:1 com users)';
COMMENT ON COLUMN members.role IS 'Cargo/função do membro: admin, tecnico, auxiliar ou estagiario';
COMMENT ON COLUMN members.is_active IS 'Indica se o membro está ativo para atribuições';
COMMENT ON COLUMN members.created_at IS 'Data e hora de criação do registro';
COMMENT ON COLUMN members.updated_at IS 'Data e hora da última atualização (trigger automático)';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS members CASCADE;
DROP TYPE IF EXISTS member_role;
-- +goose StatementEnd
