-- +goose Up
CREATE TABLE employees (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    emp_name TEXT NOT NULL,
    department TEXT NOT NULL,
    emp_role TEXT NOT NULL
);

-- +goose Down
DROP TABLE employees;
