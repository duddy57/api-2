-- name: CreateMemberQuery :exec
INSERT INTO members (user_id, role)
VALUES ($1, $2);

-- name: GetMemberQuery :many
SELECT
    m.id,
    m.role,
    u.email,
    u.username
FROM members m
JOIN users u ON m.user_id = u.id;
