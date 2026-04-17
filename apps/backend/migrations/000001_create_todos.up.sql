CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS todos (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO todos (id, title, description, completed, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'Set up backend structure', 'Bootstrap clean architecture with handlers, usecases, repositories, and migrations.', FALSE, NOW(), NOW()),
    (gen_random_uuid(), 'Connect frontend to backend', 'Fetch health and todos from Next.js using the shared API client.', FALSE, NOW(), NOW());
