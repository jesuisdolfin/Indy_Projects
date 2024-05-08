-- name: CreateUser :one
INSERT INTO user_data (
    name,
    weight,
    height,
    age
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM user_data 
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM user_data
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE user_data
SET weight = 170.0
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM user_data
WHERE id = $1;