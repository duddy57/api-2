-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Tabela: users
-- Descrição: Armazena usuários do sistema com credenciais de autenticação
-- Versão: 2.0
-- ============================================================================

CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuidv7(),

    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password_hash BYTEA NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT users_email_unique UNIQUE (email),
    CONSTRAINT users_username_unique UNIQUE (username),
    CONSTRAINT users_email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    CONSTRAINT users_username_length CHECK (char_length(username) >= 3)
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at DESC);

COMMENT ON TABLE users IS 'Usuários do sistema com credenciais de autenticação';
COMMENT ON COLUMN users.id IS 'Identificador único do usuário (UUID)';
COMMENT ON COLUMN users.username IS 'Nome de usuário único (mínimo 3 caracteres)';
COMMENT ON COLUMN users.email IS 'E-mail único do usuário (validado por regex)';
COMMENT ON COLUMN users.password_hash IS 'Hash da senha (bcrypt/argon2)';
COMMENT ON COLUMN users.created_at IS 'Data e hora de criação do registro';
COMMENT ON COLUMN users.updated_at IS 'Data e hora da última atualização (trigger automático)';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
