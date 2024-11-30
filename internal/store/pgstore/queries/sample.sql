-- name: GetAllSamples :many
SELECT * FROM "sample" ORDER BY "created_at" DESC;

-- name: GetSampleByID :one
SELECT * FROM "sample" WHERE "id" = $1;

-- name: CreateSample :one
INSERT INTO "sample" ("name", "password", "created_at") VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateSample :one
UPDATE "sample" SET "name" = $1, "password" = $2, "created_at" = $3 WHERE "id" = $4 RETURNING *;

-- name: DeleteSample :exec
DELETE FROM "sample" WHERE "id" = $1;