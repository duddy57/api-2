-- name: CreateFormTecnicoQuery :copyfrom
INSERT INTO form_tecnico (
    member_id,
    form_id
)
VALUES (
    $1,
    $2
    );

-- name: CreateFormQuery :one
INSERT INTO forms (
    client_id,
    solicited_name,
    difficulty_level,
    defect_description,
    solution_description,
    occurred_at
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetFormByIdQuery :one
SELECT
    f.id,
    f.client_id,
    c.name as client_name,
    f.occurred_at,
    f.solicited_name,
    f.difficulty_level,
    f.defect_description,
    f.solution_description,
    f.created_at,
    f.updated_at
FROM forms f
JOIN clients c ON f.client_id = c.id
WHERE f.id = $1;

-- name: GetFormsQuery :many
SELECT
    f.id,
    f.client_id,
    c.name as client_name,
    f.occurred_at,
    f.solicited_name,
    f.difficulty_level,
    f.defect_description,
    f.solution_description,
    f.created_at,
    f.updated_at
FROM forms f
JOIN clients c ON f.client_id = c.id
ORDER BY f.id ASC;

-- name: GetFormTecnicosByFormID :many
SELECT
    form_tecnico.id,
    form_tecnico.member_id,
    form_tecnico.form_id,
    form_tecnico.created_at,
    form_tecnico.updated_at,
    users.username AS user_name,
    users.email AS user_email
FROM form_tecnico
JOIN members ON form_tecnico.member_id = members.id
JOIN users ON members.user_id = users.id
WHERE form_tecnico.form_id = $1
ORDER BY form_tecnico.id ASC;


-- name: UpdateFormQuery :exec
UPDATE forms
SET client_id = $1,
    solicited_name = $2,
    difficulty_level = $3,
    defect_description = $4,
    solution_description = $5,
    updated_at = NOW()
WHERE id = $6;

-- name: DeleteFormQuery :exec
DELETE FROM forms
WHERE id = $1;
