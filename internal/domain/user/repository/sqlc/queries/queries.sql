-- name: CreateUser :one
INSERT INTO
    users (email, username, password, bio, image)
VALUES
    (?, ?, ?, ?, ?) RETURNING id;