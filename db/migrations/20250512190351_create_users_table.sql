-- +goose Up
CREATE TABLE public.users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    email VARCHAR(255) NOT NULL UNIQUE,
    full_name VARCHAR(255),
    phone VARCHAR(16),
    password VARCHAR(255),
    refresh_tokenoh TEXT
);

CREATE INDEX idx_users_deleted_at ON public.users(deleted_at);


-- +goose Down
DROP TABLE public.users;