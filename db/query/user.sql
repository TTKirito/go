-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- -- name: UpdateUser :one
-- UPDATE users 
-- SET 
--     hashed_password = CASE when @set_hashed_password::boolean = TRUE THEN @hashed_password ELSE hashed_password END,
--     full_name = CASE when @set_full_name::boolean = TRUE THEN @full_name ELSE full_name END,
--     email = CASE when @set_email::boolean = TRUE THEN @email ELSE email END
-- WHERE 
--     username = @username
-- RETURNING *;

-- name: UpdateUser :one
UPDATE users 
SET 
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email)
WHERE 
    username = sqlc.arg(username)
RETURNING *;