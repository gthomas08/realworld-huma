-- name: CreateUser :exec
INSERT INTO
    users (email, username, password, bio, image)
VALUES
    (?, ?, ?, ?, ?) RETURNING id;