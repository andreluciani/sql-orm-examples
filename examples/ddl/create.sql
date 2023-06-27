CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title varchar NOT NULL,
    content varchar NOT NULL,
    create_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);