-- name: CreateClientQuery :one
INSERT INTO clients (
  name,

  email,
  phone,
  contact_name,
  client_type,
  cnpj_cpf,

  postal_code,
  neighborhood,
  country,
  state,
  city,
  street,
  number,
  complement,
  latitude,
  longitude
   )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
RETURNING id;

-- name: UpdateClientQuery :exec
UPDATE clients
SET
  name = $1,

  email = $2,
  phone = $3,
  contact_name = $4,
  client_type = $5,
  cnpj_cpf = $6,
  postal_code = $7,
  neighborhood = $8,
  country = $9,
  state = $10,
  city = $11,
  street = $12,
  number = $13,
  complement = $14,
  latitude = $15,
  longitude = $16
WHERE id = $17;

-- name: GetAllClientsQuery :many
SELECT
  id,
  name,

  email,
  phone,
  contact_name,
  client_type,
  cnpj_cpf,

  postal_code,
  neighborhood,
  country,
  state,
  city,
  street,
  number,
  complement,
  latitude,
  longitude,

  created_at,
  updated_at
FROM clients
ORDER BY id DESC;

-- name: GetClientByIdQuery :one
SELECT
  id,
  name,
  email,
  phone,
  contact_name,
  client_type,
  cnpj_cpf,

  postal_code,
  neighborhood,
  country,
  state,
  city,
  street,
  number,
  complement,
  latitude,
  longitude,

  created_at,
  updated_at
FROM clients
WHERE id = $1;

-- name: DeleteClientQuery :exec
DELETE FROM clients
WHERE id = $1;

-- -- name: GetCalledForClienteQuery :many
-- SELECT
--   cl.id,
--   cl.public_code,

--   cl.occurrence_date,
--   cl.os_type,
--   cl.public_code,
--   cl.started_at,
--   cl.finished_at,

--   cl.defect_description,
--   cl.status,
--   ct.client_name as client_name,
--   st.name as service_name,
--   rt.team_name as responsible_team_name,
--   cl.created_at,
--   cl.updated_at
-- FROM calleds cl
-- INNER JOIN clients ct ON cl.client_id = ct.id
-- INNER JOIN services st ON cl.service_id = st.id
-- INNER JOIN teams rt ON cl.responsible_team_id = rt.id
-- WHERE cl.client_id = $1
-- ORDER BY cl.created_at DESC;
