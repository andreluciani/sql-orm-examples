-- Without a transaction:
INSERT INTO cities(name)
VALUES ('New York');
INSERT INTO users(name, email, city_id, role_id)
VALUES ('John', 'john@email.com', (SELECT city_id FROM cities WHERE name = 'New York'), 1);

-- With a transaction:
BEGIN;
INSERT INTO cities(name)
VALUES ('New York');
INSERT INTO users(name, email, city_id, role_id)
VALUES ('John', 'john@email.com', (SELECT city_id FROM cities WHERE name = 'New York'), 1);
COMMIT;