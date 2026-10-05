CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    phone_number VARCHAR(255),
    user_password VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT pk_users PRIMARY KEY (id),
    CONSTRAINT uq_users UNIQUE (email, phone_number)
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    id SERIAL,
    user_id INT REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    refresh_token TEXT,
    expiry TIMESTAMPTZ NOT NULL,
    revoked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT now(),

    CONSTRAINT pk_refresh_tokens PRIMARY KEY (id)

);