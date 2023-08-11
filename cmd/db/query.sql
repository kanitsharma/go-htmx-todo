-- name: ListTodos :many
SELECT * FROM todos 
ORDER BY name;

-- name: CreateTodo :one
INSERT INTO todos(
  name, status
) VALUES (
  ?, ?
)
RETURNING *;
