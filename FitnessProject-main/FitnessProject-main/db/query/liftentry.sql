-- name: CreateLiftEntry :one
INSERT INTO liftentries (
    user_id,
    weight_lifted,
    reps
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetLiftEntry :one
SELECT * FROM liftentries
WHERE id = $1 LIMIT 1;

-- name: ListLiftEntries :many
SELECT * FROM liftentries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateLiftEntry :one
UPDATE liftentries
SET reps = 5
WHERE id = $1
RETURNING *;

-- name: DeleteLiftEntry :exec
DELETE FROM liftentries
WHERE id = $1;