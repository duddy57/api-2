-- name: CreateUserQuery :one
INSERT INTO users (username, email, password_hash, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;


-- name: GetUserByEmailQuery :one
SELECT id, username, email, password_hash
FROM users
WHERE email = $1;

-- name: GetUserByIdQuery :one
SELECT u.id, u.username, u.email, m.role, u.created_at, u.updated_at
FROM users u
JOIN members m ON u.id = m.user_id
WHERE u.id = $1;


-- -- name: GetTechs :many
-- SELECT id, user_name, email, role, created_at, updated_at
-- FROM users
-- WHERE role = 'tecnico';

-- name: UpdateUserQuery :exec
UPDATE users
SET username = $1, updated_at = now()
WHERE id = $2;

-- name: DeleteUserQuery :exec
DELETE FROM users
WHERE id = $1;

-- -- name: GetClients :many
-- SELECT id, user_name, email, role, created_at, updated_at
-- FROM users
-- WHERE role = 'CLIENT';

-- -- name: GetServices :many
-- SELECT id, user_name, email, role, created_at, updated_at
-- FROM users
-- WHERE role = 'SERVICE';
