CREATE TABLE users (
    id serial PRIMARY KEY,
    first_name VARCHAR(50)
);
CREATE TABLE salaries (user_id int UNIQUE NOT NULL, amount int);
ALTER TABLE salaries
ADD CONSTRAINT users_salaries_fk0 FOREIGN KEY (user_id) REFERENCES users (id);