DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar NOT NULL,
    email varchar NOT NULL UNIQUE,
    city_id integer NOT NULL,
    role_id integer NOT NULL
);
ALTER TABLE users
ADD CONSTRAINT users_fk0 FOREIGN KEY (city_id) REFERENCES cities(id);
ALTER TABLE users
ADD CONSTRAINT users_fk1 FOREIGN KEY (role_id) REFERENCES roles(id);
INSERT INTO users(name, email, city_id, role_id)
VALUES (
        'Viviana',
        'viviana@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Callan',
        'callan@email.com',
        ceil(random() * 5),
        2
    ),
    (
        'Aila',
        'aila@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Moses',
        'moses@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Amelia',
        'amelia@email.com',
        ceil(random() * 5),
        2
    ),
    (
        'Chandler',
        'chandler@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Alicia',
        'alicia@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Nehemiah',
        'nehemiah@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Everly',
        'everly@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Kayson',
        'kayson@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Imani',
        'imani@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Jamie',
        'jamie@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Ximena',
        'ximena@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Alexis',
        'alexis@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Estrella',
        'estrella@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Talon',
        'talon@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Aya',
        'aya@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Finnegan',
        'finnegan@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Lylah',
        'lylah@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Cooper',
        'cooper@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Audrey',
        'audrey@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Princeton',
        'princeton@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Paislee',
        'paislee@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Arjun',
        'arjun@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Scarlett',
        'scarlett@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Andrew',
        'andrew@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Charlie',
        'charlie@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Emanuel',
        'emanuel@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Molly',
        'molly@email.com',
        ceil(random() * 5),
        1
    ),
    (
        'Riley',
        'riley@email.com',
        ceil(random() * 5),
        1
    );