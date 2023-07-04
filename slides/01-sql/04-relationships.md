---
marp: true
title: SQL and ORM
description: Illustrating fundamentals of SQL and ORM
theme: base
transition: fade
paginate: false
_paginate: false
math: mathjax
---

<!-- _class: invert -->
<!-- _transition: cover -->

# Database Relationships

There are three main SQL database relationships:

1. One-to-one
2. One-to-many (or Many-to-one)
3. Many-to-many

**Let's see when to use each of them!**

---

# _One-to-One_

---

<!-- _class: invert -->

# _One-to-One_

This relation is used when **one** row of `table_one` is linked (or related) to only **one** row of `table_two`

Example: A person and their birthplace, an employee and their salary, a user and their role

---

<!-- _class: invert -->

# _One-to-One_

![h:500](/assets/one-to-one.png)

---

<!-- _class: invert -->

# _One-to-One_

To achieve a one-to-one relationship between tables:

1. Create the two tables that will be linked
2. On one table, add a column with the `FOREIGN KEY` constraint.
3. The `FOREIGN KEY` column **must be** `UNIQUE`

---

<!-- _class: invert -->
<!-- _transition: cover -->

# _One-to-One_

```sql
CREATE TABLE users (
  id serial PRIMARY KEY,
  first_name VARCHAR(50)
);

CREATE TABLE salaries (
  user_id int UNIQUE NOT NULL,
  amount int
);

ALTER TABLE salaries
ADD CONSTRAINT users_salaries_fk0
FOREIGN KEY (user_id)
REFERENCES users (id);
```

---

# _One-to-Many_

---

<!-- _class: invert -->

# _One-to-Many_

The _one-to-many_ relation is the most used relation in SQL and occurs when **one** row from `table_one` is related one or **many** rows in `table_two`.

Example: A country and its states or cities, students and the class they belong to, employees and their department, among several others.

---

<!-- _class: invert -->

# _One-to-Many_

![h:200](/assets/one-to-many.png)

---

<!-- _class: invert -->

# _One-to-Many_

To achieve a one-to-many relationship between tables:

1. Create the two tables that will be linked
2. On one table, add a column with the `FOREIGN KEY` constraint.

---

<!-- _class: invert -->
<!-- _transition: cover -->

# _One-to-Many_

```sql
CREATE TABLE countries (
    id serial PRIMARY KEY,
    name varchar NOT NULL,
);
CREATE TABLE cities (
    id serial PRIMARY KEY,
    country_id integer NOT NULL,
    name varchar NOT NULL,
);
ALTER TABLE cities
ADD CONSTRAINT cities_fk0
FOREIGN KEY (country_id)
REFERENCES countries(id);
```

---

## _One-to-Many_ or _Many-to-One_ ?

- While many people do not differentiate the two relations, it is just a matter of focus.

- If we take the example shown: countries and cities. A country is consisted of many cities, in this case, coutry to city is a _one-to-many_ relationship.

- On the other hand, if we focus on the cities, we can say that many cities are part of one country, resulting in a _many-to-one_ relationship.

---

# _Many-to-Many_

---

<!-- _class: invert -->

# _Many-to-Many_

The _many-to-many_ relation occurs when **many** rows from `table_one` are related to **many** rows in `table_two`.

Example: Product and suppliers, flights and passengers, etc.

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 22pt;
}
</style>

# _Many-to-Many_

![h:350](/assets/many-to-many.png)

- _Many-to-many_ relation between `customers` and `flights` throught the `passengers` table
- The `flights` table has two _one-to-one_ relationships with `airports` table.

---

<!-- _class: invert -->

# _Many-to-Many_

To achieve a many-to-many relationship between tables:

1. Create the two tables that will be linked
2. Create a third table (called _linking_, _bridging_ or _junction_ table)
3. The third table will store the primary keys of both the tables to be linked

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 90%;
}
</style>

# _Many-to-Many_

Creating the tables:

```sql
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
CREATE TABLE flights (
    id serial PRIMARY KEY,
    from_airport_id integer NOT NULL,
    to_airport_id integer NOT NULL,
);
```

---

<!-- _class: invert -->
<!-- _transition: cover -->

# _Many-to-Many_

Adding the constraints:

```sql
ALTER TABLE flights
ADD CONSTRAINT flights_fk0
FOREIGN KEY (from_airport_id)
REFERENCES airports(id);

ALTER TABLE flights
ADD CONSTRAINT flights_fk1
FOREIGN KEY (to_airport_id)
REFERENCES airports(id);
```
