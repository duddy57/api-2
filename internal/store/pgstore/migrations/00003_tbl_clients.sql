-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Tabela: clients
-- Descrição: Clientes do sistema (empresas e pessoas físicas)
-- Versão: 2.0
-- ============================================================================

CREATE TYPE client_type AS ENUM('avulso', 'contrato');
CREATE TABLE IF NOT EXISTS clients (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    name VARCHAR(100) NOT NULL,
    client_type client_type NOT NULL DEFAULT 'avulso',
    cnpj_cpf VARCHAR(18),

    email VARCHAR(100),
    phone VARCHAR(20),
    contact_name VARCHAR(100),

    street VARCHAR(255),
    number VARCHAR(10),
    neighborhood VARCHAR(100),
    city VARCHAR(100),
    state VARCHAR(2),
    country VARCHAR(2) DEFAULT 'BR',
    postal_code VARCHAR(10),

    complement VARCHAR(255),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT clients_email_format CHECK (email IS NULL OR email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    CONSTRAINT clients_name_length CHECK (char_length(name) >= 3),
    CONSTRAINT clients_cnpj_cpf_length CHECK (cnpj_cpf IS NULL OR char_length(cnpj_cpf) >= 11),
    CONSTRAINT clients_latitude_range CHECK (latitude IS NULL OR (latitude >= -90 AND latitude <= 90)),
    CONSTRAINT clients_longitude_range CHECK (longitude IS NULL OR (longitude >= -180 AND longitude <= 180))
);

CREATE INDEX IF NOT EXISTS idx_clients_phone ON clients(phone) WHERE phone IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_clients_cnpj_cpf ON clients(cnpj_cpf) WHERE cnpj_cpf IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_clients_type ON clients(client_type);
CREATE INDEX IF NOT EXISTS idx_clients_city ON clients(city) WHERE city IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_clients_state ON clients(state) WHERE state IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_clients_created_at ON clients(created_at DESC);

CREATE INDEX IF NOT EXISTS idx_clients_location ON clients(latitude, longitude)
    WHERE latitude IS NOT NULL AND longitude IS NOT NULL;


COMMENT ON TABLE clients IS 'Clientes do sistema (empresas e pessoas físicas) - avulso ou contrato';
COMMENT ON COLUMN clients.id IS 'Identificador único do cliente (UUID)';
COMMENT ON COLUMN clients.name IS 'Nome/Razão social do cliente (mínimo 3 caracteres)';
COMMENT ON COLUMN clients.email IS 'E-mail do cliente (opcional, validado por regex)';
COMMENT ON COLUMN clients.phone IS 'Telefone de contato';
COMMENT ON COLUMN clients.contact_name IS 'Nome da pessoa de contato';
COMMENT ON COLUMN clients.client_type IS 'Tipo de cliente: avulso (serviços pontuais) ou contrato (recorrente)';
COMMENT ON COLUMN clients.cnpj_cpf IS 'CNPJ (18 chars) ou CPF (14 chars com formatação)';
COMMENT ON COLUMN clients.postal_code IS 'CEP do endereço';
COMMENT ON COLUMN clients.neighborhood IS 'Bairro';
COMMENT ON COLUMN clients.country IS 'Código do país (ISO 3166-1 alpha-2) - padrão BR';
COMMENT ON COLUMN clients.state IS 'UF (sigla do estado)';
COMMENT ON COLUMN clients.city IS 'Cidade';
COMMENT ON COLUMN clients.street IS 'Logradouro (rua, avenida, etc)';
COMMENT ON COLUMN clients.number IS 'Número do endereço';
COMMENT ON COLUMN clients.complement IS 'Complemento do endereço';
COMMENT ON COLUMN clients.latitude IS 'Latitude para geolocalização (-90 a 90)';
COMMENT ON COLUMN clients.longitude IS 'Longitude para geolocalização (-180 a 180)';
COMMENT ON COLUMN clients.created_at IS 'Data e hora de criação do registro';
COMMENT ON COLUMN clients.updated_at IS 'Data e hora da última atualização (trigger automático)';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients CASCADE;
DROP TYPE IF EXISTS client_type;
-- +goose StatementEnd
