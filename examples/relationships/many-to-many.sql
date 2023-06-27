CREATE TABLE airports (
    id serial PRIMARY KEY,
    name varchar NOT NULL,
    code varchar NOT NULL,
    city_id integer NOT NULL,
);
CREATE TABLE flights (
    id serial PRIMARY KEY,
    from_airport_id integer NOT NULL,
    to_airport_id integer NOT NULL,
);
CREATE TABLE customers (
    id serial PRIMARY KEY,
    email varchar NOT NULL UNIQUE,
    name varchar NOT NULL,
);
CREATE TABLE passengers (
    id serial PRIMARY KEY,
    customer_id integer NOT NULL,
    flight_id integer NOT NULL,
);
ALTER TABLE flights
ADD CONSTRAINT flights_fk0 FOREIGN KEY (from_airport_id) REFERENCES airports(id);
ALTER TABLE flights
ADD CONSTRAINT flights_fk1 FOREIGN KEY (to_airport_id) REFERENCES airports(id);
ALTER TABLE passengers
ADD CONSTRAINT passengers_fk0 FOREIGN KEY (flight_id) REFERENCES flights(id);
ALTER TABLE passengers
ADD CONSTRAINT passengers_fk1 FOREIGN KEY (flight_id) REFERENCES flights(id);