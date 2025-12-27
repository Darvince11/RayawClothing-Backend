CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    user_password VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    refresh_token TEXT,
    user_id INT REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);