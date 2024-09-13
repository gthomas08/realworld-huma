-- name: GetPingMessage :one
SELECT message FROM pings
LIMIT 1;
