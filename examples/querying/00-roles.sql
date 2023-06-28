DROP TABLE IF EXISTS roles CASCADE;
CREATE TABLE roles (id serial PRIMARY KEY, name varchar NOT NULL);
INSERT INTO roles(name)
VALUES ('admin'),
    ('basic');