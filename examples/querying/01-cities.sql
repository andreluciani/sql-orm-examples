DROP TABLE IF EXISTS cities CASCADE;
CREATE TABLE cities (
    id serial PRIMARY KEY,
    name varchar NOT NULL,
    population integer NOT NULL
);
INSERT INTO cities(name, population)
VALUES ('São Paulo', 12396372),
    ('Rio de Janeiro', 6775561),
    ('Brasília', 3094325),
    ('Salvador', 2900319),
    ('Fortaleza', 2703391);