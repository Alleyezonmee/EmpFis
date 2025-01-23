-- name: CreateEmployee :one
INSERT INTO employees (id, created_at, updated_at, emp_role, emp_name, department)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;