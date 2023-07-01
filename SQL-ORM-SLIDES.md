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

<!-- _class: lead -->

# <!--fit--> Fundamentals of SQL and ORM

Understanding SQL at its Core: Building a Strong Foundation for Data Management

<!-- The aim goal of this presentation is to show the main concepts of SQL and ORMs. -->

---

<!-- _class: invert -->

# Agenda

- ![1](https://icongr.am/material/numeric-1-circle.svg?color=ffffff) **SQL**
  - Definition, basic syntax, querying and more
- ![2](https://icongr.am/material/numeric-2-circle.svg?color=ffffff) **ORMs**
  - Definition, popular ORMs and examples
- ![3](https://icongr.am/material/numeric-3-circle.svg?color=ffffff) **Best Practices and Tips**
  - Performance, security and debugging

<!-- Here are the topics that will be covered. We'll start by looking at definitions of SQL and ORMs, including examples, followed by best practives and tips. -->

---

![1 w:128 h:128](https://icongr.am/material/numeric-1-circle.svg?color=666666)

# SQL

**S**tructured **Q**uery **L**anguage

<!-- So what is SQL? It stands for structured query language. -->

---

<style scoped>
section li em {
  font-size: 25px;
}
</style>

# ![1](https://icongr.am/material/numeric-1-circle.svg?color=666666) SQL

- It is a standard language used for managing **relational databases**
- _A relational database is a type of database that stores and provides access to data points that are related to one another. (Oracle)_
- SQL provides a set of **commands** for interacting with **databases**
- _A database is an organized collection of structured information, or data, typically stored electronically in a computer system. (Oracle)_

 <!-- It is the standard used for ralational databases, which the data points are related to one another. It provides commands to create, retrieve, update and delete (CRUD) data from and to the database. -->

---

### SQL Database Management Systems

|                                                        DMS                                                         |         License         |
| :----------------------------------------------------------------------------------------------------------------: | :---------------------: |
|          ![MySQL Logo](https://icongr.am/simple/mysql.svg?size=64&color=currentColor&colored=false) MySQL          | Proprietary/Open-source |
| ![](https://icongr.am/simple/microsoftsqlserver.svg?size=64&color=currentColor&colored=false) Microsoft SQL Server |       Proprietary       |
|              ![](https://icongr.am/simple/oracle.svg?size=64&color=currentColor&colored=false) Oracle              |       Proprietary       |
|          ![](https://icongr.am/simple/postgresql.svg?size=64&color=currentColor&colored=false) PostgreSQL          |       Open-source       |
|              ![](https://icongr.am/simple/sqlite.svg?size=64&color=currentColor&colored=false) SQLite              |       Open-source       |

 <!-- Here are a few examples of Relational Database Management Systems that use SQL -->

---

<!-- _class: invert -->

### SQL Database Example

![h:550](./assets/example-database.png)

 <!-- Take a look at this database architecture. It shows tables and relations between them. In the next slides we'll understand how everything is linked and model ourselves a few examples. -->

---

<!-- _class: invert -->

### Databases, Tables, Rows and Columns

Imagine we have a PostgreSQL server running:

![postgres](https://icongr.am/simple/postgresql.svg?size=128&color=ffffff&colored=false)
_postgres://localhost:5432_

- #### Databases

- `blog`
- `supermarket`
- `gym`

 <!-- In this server, we can have multiple databases with different purposes and context, for instance, a database for a blog, a database for a supermarket and a database for a gym. Each of them can store data that is relevant for instance, users and blog posts for the first one, products and prices for the second and users and workout plans for the latest -->

---

<!-- _class: invert -->

### Databases, Tables, Rows and Columns

<div class="columns">
<div>

Databases act like _containers_ of related data.

In relational databases, the data is stored in **tables**

</div>
<div>

`blog` database:

![h:190](./assets/tables.png)

4 tables, 3 _relations_

</div>
</div>

 <!-- In this example, the blog database has 4 tables, one for storing user data, one for storing posts data, one for storing user types (basic, admin or superuser, for instance) and one to map users and posts, which is called a junction table. These links are called relations and there are some types of relations possibles, which will be shown later. -->

---

<!-- _class: invert -->

### Databases, Tables, Rows and Columns

#### Table _Columns_

- `users`
  - `id` - int - **Primary** key
  - `user_type` - int - **Foreign** key
  - `name` - varchar
  - `...`
  - `created_at` - datetime

 <!-- Each table has columns, which are attributes of every entry (row) of the table. You can think of tables as spreadsheets. Each column has a type (int, char, datetime, enum, etc), and can have a constraint as well (unique, not null, etc). If the column is a **primary** key, it means that the column identifies the row of the table. If the column is a **foreign** key, that means it is related to ANOTHER table, and it uses the primary key OR another unique column of the other table. -->

---

<!-- _class: invert -->

### Databases, Tables, Rows and Columns

#### Table _Rows_

```bash {4}
blog=# select id, user_type_id, name, email, created_at from users;
 id | user_type_id |   name    |           email           |         created_at
----+--------------+-----------+---------------------------+----------------------------
  1 |            1 | André     | andre.luciani@email.com   | 2023-06-25 20:31:46.438607
  2 |            1 | John      | john.doe@email.com        | 2023-06-25 20:31:46.438607
  3 |            1 | Priscilla | priscilla.scott@email.com | 2023-06-25 20:31:46.438607
(3 rows)
```

 <!-- In this slide we can see some ROWS from the "users" table. Each row is one entry and has the attributes defined in the columns. -->

---

### SQL's Role in Database Management

- SQL is specifically designed for managing **relational** databases.

- It provides a standardized approach for _creating_, _modifying_, and _querying_ data.

- SQL ensures data _integrity_ and _consistency_ in relational databases.

 <!-- In summary, SQL is used to manage relational databases while ensuring data integrity and consistency. -->

---

  <!-- _class: invert -->

### Data Manipulation Language (DML)

- Subset of a programming language used explicitly to make changes in the database
  _(e.g. CRUD operations)_.

- In case of SQL:

  - `SELECT` - Retrieve

  - `INSERT` - Create

  - `UPDATE` - Modify existing data

  - `DELETE` - Exclude entries

 <!-- Data manipulation language is a subset of a programming language used to make changes in the database. Let's see some of them in the following slides. -->

---

  <!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-sql {
  font-size: 150%;
}
</style>

### The `SELECT` statement

- Used to retrieve data. Basic syntax:

```sql
SELECT
  column1, column2, column3
FROM
  table_name;
```

- The wildcard `*` can be used to select "all columns"

- When making SQL statementes, always remember to close the statement with `;`

- The SQL words (`SELECT`, `FROM`, ...) usually are written with uppercase for better readability.

 <!-- The select statement is one of the most important. It is used to retrieve data and the basic syntax is as follows. -->

---

  <!-- _class: invert -->

### The `SELECT` statement

###### Example using [`psql`](https://www.postgresql.org/docs/current/app-psql.html):

```bash
psql      # starts postgres interactive terminal
\c blog   # connects to 'blog' database
```

###### After connecting to the database, execute the query:

```
blog=# SELECT id, name, last_name, email FROM users;
 id |   name    | last_name |           email
----+-----------+-----------+---------------------------
  1 | André     | Luciani   | andre.luciani@email.com
  2 | John      | Doe       | john.doe@email.com
  3 | Priscilla | Scott     | priscilla.scott@email.com
(3 rows)
```

 <!-- Here's one example using psql. -->

---

  <!-- _class: invert -->

<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `INSERT` statement

- Used to add data. Basic syntax:

```sql
INSERT INTO table_name(column1, column2, …)
VALUES (value1, value2, …);
```

- Multiple rows can be added by providing more values grouped with `()`

---

  <!-- _class: invert -->

<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `INSERT` statement

##### Adding a new row on the `posts` table:

```sql
INSERT INTO posts(title, content, create_at, updated_at)
VALUES ('Another post', 'Another example', NOW(), NOW());
```

- `NOW()` is a SQL [_function_](<https://www.postgresql.org/docs/devel/functions-datetime.html#:~:text=15%3A23.5%2B01-,now%20(%20),-%E2%86%92%20timestamp%20with>) that returns the current date and time. These functions may differ depending on the DMS used.

---

  <!-- _class: invert -->

### The `INSERT` statement

```diff
id |             title              |         content          |         create_at          |         updated_at
----+--------------------------------+--------------------------+----------------------------+----------------------------
  1 | PostgreSQL 101                 | This is an example post. | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
  2 | Bread Recipe                   | This is an example post. | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
  3 | Will AI take over the world?   | This is an example post. | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
  4 | How to learn a new technology. | This is an example post. | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
+ 5 | Another post                   | Another example          | 2023-06-25 22:23:57.567231 | 2023-06-25 22:23:57.567231
(5 rows)
```

---

  <!-- _class: invert -->

<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `UPDATE` statement

- Used to update entries.
  Basic syntax:

```sql
UPDATE table_name
SET column1 = value1,
    column2 = value2,
    ...
WHERE condition;
```

<!-- When updating a table, besides telling which table and columns we want to update, we must provide a condition to filter out only the rows that we want to update too. This can be specific enough to update only one single row, or more general, if we want to update multiple entries. -->

---

  <!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `UPDATE` statement

###### Updating a row on the `posts` table:

```sql
UPDATE posts
SET content = 'The post content was updated!',
    updated_at = NOW(),
WHERE id = 2;
```

---

  <!-- _class: invert -->

### The `UPDATE` statement

```diff
 id |             title              |            content            |         create_at          |         updated_at
----+--------------------------------+-------------------------------+----------------------------+----------------------------
  1 | PostgreSQL 101                 | This is an example post.      | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
- 2 | Bread Recipe                   | This is an example post.      | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
  3 | Will AI take over the world?   | This is an example post.      | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
  4 | How to learn a new technology. | This is an example post.      | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
  5 | Another post                   | Another example               | 2023-06-25 22:23:57.567231 | 2023-06-25 22:23:57.567231
+ 2 | Bread Recipe                   | The post content was updated! | 2023-06-25 20:30:40.617806 | 2023-06-25 22:36:44.371102
```

---

  <!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `DELETE` statement

Basic syntax:

```sql
DELETE FROM table_name
WHERE condition;
```

---

  <!-- _class: invert -->

<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `DELETE` statement

###### Deleting a row from the `posts` table:

```sql
DELETE FROM posts
WHERE id = 5;
```

---

  <!-- _class: invert -->

### The `DELETE` statement

```diff
 id |             title              |            content            |         create_at          |         updated_at
----+--------------------------------+-------------------------------+----------------------------+----------------------------
  1 | PostgreSQL 101                 | This is an example post.      | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
  3 | Will AI take over the world?   | This is an example post.      | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
  4 | How to learn a new technology. | This is an example post.      | 2023-06-25 20:30:40.617806 | 2023-06-25 20:30:40.617806
- 5 | Another post                   | Another example               | 2023-06-25 22:23:57.567231 | 2023-06-25 22:23:57.567231
  2 | Bread Recipe                   | The post content was updated! | 2023-06-25 20:30:40.617806 | 2023-06-25 22:36:44.371102
```

---

  <!-- _class: invert -->

### Data Definition Language (DDL)

- Subset of a programming language used to _define_ or _describe_ databases **schemas** (the "structure" of the DB).

- A few examples from SQL:

  - `CREATE TABLE` - Create tables

  - `ALTER TABLE` - Modify existing tables

  - `DROP TABLE` - Exclude tables

 <!-- If DML is used to CRUD entries, the DDL is used to CRUD structures. For instance, if you want to add a new column to an existing table, you will use DML. -->

---

  <!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-sql {
  font-size: 120%;
}
</style>

### The `CREATE TABLE` statement

- Used to create tables. Basic syntax:

```sql
CREATE TABLE table_name (
   column1 datatype(length) column_constraint,
   column2 datatype(length) column_constraint,
   column3 datatype(length) column_constraint,
   table_constraints
);
```

- Each column has a [**datatype**](https://www.postgresql.org/docs/current/datatype.html) and may have [**constraints**](https://www.postgresql.org/docs/15/ddl-constraints.html).

- The statement `CREATE TABLE IF NOT EXISTS ...` can be used to prevent an error if the table name is already in use.

 <!-- When creating a table, we must provide the name of the table, the name of each column along with its datatype and possible constraints as well as table constraints. -->

---

  <!-- _class: invert -->
<style scoped>
li {
  font-size: 30px;
  padding-left: 1px;
}
</style>

### The `CREATE TABLE` statement

#### Constraints

- Used to ensure data **consistency**

- Common constraints used in columns:

  - `NOT NULL` - values cannot be `null`

  - `UNIQUE` - all rows must have different values for the column

  - `PRIMARY KEY` - Column that is unique, not null and used to identify each row of a table.

  - `CHECK` - Boolean verifier to prevent wrong inputs
    (e.g. `CHECK (price > 0)`)

 <!-- The costraints are a powerful way to prevent wrong inputs. They are a set of rules that must be met when modifying data in a database. -->

---

  <!-- _class: invert -->
<style scoped>
li {
  font-size: 30px;
  padding-left: 1px;
}
</style>

### The `CREATE TABLE` statement

#### Constraints

- Common constraints used in columns (a few more):

  - `FOREIGN KEY` - Link tables by storing the `PRIMARY KEY`of another table.

  - `DEFAULT` - Sets a default value for a column if no value is specified

  - `CREATE INDEX` - Used to create and retrieve data from the database very quickly (more on that later)

 <!-- The costraints are a powerful way to prevent wrong inputs. They are a set of rules that must be met when modifying data in a database. -->

---

  <!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-sql {
  font-size: 120%;
}
</style>

### The `CREATE TABLE` statement

```sql
CREATE TABLE posts (
    id serial PRIMARY KEY,
    title varchar NOT NULL,
    content varchar NOT NULL,
    create_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

- The title and content could have a minimum lenght and the `create_at` and `updated_at` columns could have default values.

- The `create_at` column has a typo, thankfully not everything is doomed yet...

 <!-- Here's one example using psql. -->

---

  <!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-sql {
  font-size: 120%;
}
</style>

### The `ALTER TABLE` statement

- Used to modify tables. Usually **add**, **remove**, **rename** columns or its constraints:

```sql
ALTER TABLE table_name
RENAME COLUMN column_name
TO new_column_name;
```

```sql
ALTER TABLE table_name
ADD COLUMN column_name datatype column_constraint;
```

```sql
ALTER TABLE table_name
DROP COLUMN column_name;
```

 <!-- Here are a few examples of the ALTER TABLE syntax. -->

---

  <!-- _class: invert -->

<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `ALTER TABLE` statement

##### Fixing the typo on the `create_at` column:

```sql
ALTER TABLE posts
RENAME create_at
TO created_at;
```

---

  <!-- _class: invert -->

<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `ALTER TABLE` statement

- It is possible to add only constraints as well, for instance:

```sql
ALTER TABLE users
ADD CONSTRAINT users_fk0 FOREIGN KEY (user_type_id) REFERENCES user_types(id);
ALTER TABLE user_post
ADD CONSTRAINT user_post_fk0 FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE user_post
ADD CONSTRAINT user_post_fk1 FOREIGN KEY (post_id) REFERENCES posts(id);
```

---

  <!-- _class: invert -->

### The `DROP TABLE` statement

- As the database evolves, it might be necessary to exclude tables (e.g. a table was deprecated because the data it held is now being stored in another place)

- The `DROP TABLE` command is used for these cases.

---

  <!-- _class: invert -->

<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

### The `DROP TABLE` statement

Basic syntax:

```sql
DROP TABLE table_name;
```

---

  <!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

### The `DROP TABLE` statement

###### Deleting the `posts` table:

```sql
DROP TABLE posts;
```

###### But...

```
blog=# DROP TABLE posts;
ERROR:  cannot drop table posts because other objects depend on it
DETAIL:  constraint user_post_fk1 on table user_post depends on table posts
HINT:  Use DROP ... CASCADE to drop the dependent objects too.
```

- Since it has constraints, the operation cannot be done...

---

  <!-- _class: invert -->
  <!-- transition: cover -->

<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

### The `DROP TABLE` statement

- Using the [`CASCADE`](https://www.postgresql.org/docs/15/sql-droptable.html#:~:text=table%20to%20drop.-,CASCADE,-Automatically%20drop%20objects) solves the problem

```sql
DROP TABLE posts CASCADE;
```

```
blog=# DROP TABLE posts CASCADE;
NOTICE:  drop cascades to constraint user_post_fk1 on table user_post
DROP TABLE
```

- The operation returns which constraints were dropped because of `CASCADE`.

---

# A few considerations

---

- While modifying one column from one table is straightforward, even that can have a great impact in a large database.

- When making such modifications, we have to address how to deal with the "old" data. In some cases the solution involves setting a default value or maybe running a script to make the old data consistent with the changes introduced.

---

<!-- transition: fade -->

- The versioning of databases introduces the concept of **migrations**

- It uses scripts (that can be written in several programming languages) to make the modifications more consistent and easy to roll-back. More on that on the _ORM_ part.

- In any case, it is always recomended to make database backups before making schema changes.

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

![h:500](./assets/one-to-one.png)

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

![h:200](./assets/one-to-many.png)

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

![h:350](./assets/many-to-many.png)

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

---

# Querying with SQL

---

<!-- _class: invert -->

## Querying with SQL

- So far, we have focused on _changing_ SQL databases, both the data itself (using DML) and the "structure" schema (using DDL), defining tables, columns and relations.

- More often than not, though, databases are _read_ rather than changed

- Because of that, SQL has several features to **query** the data and that's what we will take a look at now...

---

<!-- _class: invert -->

## Querying with SQL

- The `SELECT` statement was the first DML that was presented, and it is the command we will use to make more complex queries.

- We already know how to specify which _columns_ and which _table_ we want to get results, but that is pretty limited

---

<!-- _class: invert -->

## Querying with SQL

- What if we want to filter the results based on _relations_ between two (or more) tables?

- What about _ordering_ the results based on some criteria?

- How can we generate _analytics_ from a database?

---

<!-- _class: invert -->

## Querying with SQL

- To answer all those questions, let's use a database as an example.
- There's a `start.sh` script that can be used in [Gitpod](https://gitpod.io/#https://github.com/andreluciani/sql-orm-examples) to create a PostgreSQL database exactly like the one used in the next slides.
- [PSQL](https://www.postgresql.org/docs/current/app-psql.html) was used in the slides, but any other tool works just as fine ([PgAdmin](https://www.pgadmin.org/), DataGrip, etc)

---

<!-- _class: invert -->

## Querying with SQL

- Here's the database schema diagram used in the examples:
  ![w:1000](./assets/blog-schema.png)
- It is a blog database with six tables

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

## The `LIMIT` clause

- To limit the results returned, we can append the `LIMIT` keyword in a query, followed by the maximum number of results expected:

```sql
SELECT * FROM table_name LIMIT 5;
```

- When in doubt of the size of a table, it is a good practice to put a limit in the results to prevent slow queries (imagine querying a table with _thousands_ of rows 😅)

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

## The `LIMIT` clause

- Seeing the query in practice:

```sql
blog=# SELECT * FROM users LIMIT 5;
 id |  name   |       email       | city_id | role_id
----+---------+-------------------+---------+---------
  1 | Viviana | viviana@email.com |       5 |       1
  2 | Callan  | callan@email.com  |       2 |       2
  3 | Aila    | aila@email.com    |       5 |       1
  4 | Moses   | moses@email.com   |       2 |       1
  5 | Amelia  | amelia@email.com  |       3 |       2
(5 rows)
```

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

## The `WHERE` clause

- Another way we can refine SQL queries is using the `WHERE` keyword:

```sql
SELECT * FROM table_name
WHERE condition;
```

- The `condition` can vary a lot. Some examples are:
  - a numeric column is greater than some value
  - a column is not null
  - the row was created before a date
  - and so on...

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 140%;
}
</style>

## The `WHERE` clause

- In the `LIMIT` example, there was a column `role_id` with different values (1 and 2). Let's filter only the rows where the column `role_id` is equal to 2:

```sql
SELECT * FROM users
WHERE role_id=2;
```

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

## The `WHERE` clause

```sql
blog=# SELECT * FROM users WHERE role_id=2;
 id |  name  |      email       | city_id | role_id
----+--------+------------------+---------+---------
  2 | Callan | callan@email.com |       2 |       2
  5 | Amelia | amelia@email.com |       3 |       2
(2 rows)
```

- We can see there are two users with the `role_id` equal to 2

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

## The `ORDER BY` clause

- If we want to specify how the results should be organized, we can use the `ORDER BY` statement:

```sql
SELECT * FROM table_name
ORDER BY column ASC;
```

- The `ASC` (default) keyword orders the results in **asc**ending order
- The `DESC` keyword orders the result in **desc**ending order

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 140%;
}
</style>

## The `ORDER BY` clause

- Let's order the cities stored in table `cities` based on the `population` colum, from the most inhabited to the least inhabited:

```sql
SELECT * FROM cities
ORDER BY population DESC;
```

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

## The `ORDER BY` clause

```sql
blog=# SELECT * FROM cities ORDER BY population DESC;
 id |      name      | population
----+----------------+------------
  1 | São Paulo      |   12396372
  2 | Rio de Janeiro |    6775561
  3 | Brasília       |    3094325
  4 | Salvador       |    2900319
  5 | Fortaleza      |    2703391
(5 rows)
```

---

<style scoped>
li {
  font-size: 80%;
}
</style>

# Joining Tables

- We have seen how to filter results with `WHERE` and ordering with `ORDER`, and we also know how to create _relations_ using `FOREIGN KEY`s.

- But until now, we only get results from a single table 😔
- That makes it hard to truly understand some data, for example:
  - What is the actual role of users with `role_id=2`?
  - Who is the author of posts with `user_id=5`?
  - What are the cities with the highest number of blog posts?

---

# Joining Tables

- To answer all those questions we are going to use the `JOIN` statement.

- SQL joins allow us to combine two or more tables based on a condition, which usually is a pair of columns that are equal on the tables being joined (_equi JOINs_)

- It **is** possible to join with non-equal conditions (_non-equi JOINs_)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN`

- The `LEFT JOIN` clause returns _all_ the values from the "left" table and also the values from the "right" table where the _join condition_ is met.

- Let's use a Venn diagram to visualize it better...

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN`

![h:400](./assets/left-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN`

- The basic syntax is:

```sql
SELECT table_one.column_name, table_two.column_name
FROM table_one
LEFT JOIN table_two
ON condition;
```

 <!-- We have to define which columns to be returned, but since there might be columns with the same name in different tables, we also need to specify the table name along with the column name. The condition may vary but usually is a equality between the columns used to create the relationships -->

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN`

- Cheking what is the actual role of the users:

```sql
SELECT users.name, roles.name
FROM users
LEFT JOIN roles
ON users.role_id = roles.id
LIMIT 5;
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN`

- Cheking what is the actual role of the users:

```sql
  name   | name
---------+-------
 Viviana | admin
 Callan  | basic
 Aila    | admin
 Moses   | admin
 Amelia  | basic
(5 rows)
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN`

- To improve readability or simplify, we can use the `AS` clause to rename both the elements of the query and the columns returned:

```sql
SELECT u.name, r.name AS role
FROM users AS u
LEFT JOIN roles AS r
ON u.role_id = r.id
LIMIT 5;
```

---

<!-- _class: invert -->
<style scoped>
  li {
    font-size: 70%;
  }
  </style>

# The `JOIN` statemnet

## `LEFT JOIN`

- To improve readability or simplify, we can use the `AS` clause to rename both the elements of the query and the columns returned:

```sql
  name   | role
---------+-------
 Viviana | admin
 Callan  | basic
 Aila    | admin
 Moses   | admin
 Amelia  | basic
(5 rows)
```

---

<style scoped>
  li {
    font-size: 80%;
  }
</style>

## The `JOIN` statemnet

### `LEFT JOIN`

- In the example used, the results were limited to 5 rows using the `LIMIT 5` statement.

- But even if the results were not limited, all the users would have a matching role (you can check it yourself 😉).

- If there were a user **without** a `role_id` defined (not possible because of the constraint `NOT NULL`), then the `role` column would be `NULL`

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN` (excluding)

- The excluding `LEFT JOIN` returns all the values from the left table that **do not have** a matching row in the right table.

- Let's use a Venn diagram once again...

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN` (excluding)

![h:400](./assets/left-excluding-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN` (excluding)

- The basic syntax is:

```sql
SELECT table_one.column_name, table_two.column_name
FROM table_one
LEFT JOIN table_two
ON table_one.column_name = table_two.column_name;
WHERE table_two.column_name IS NULL;
```

 <!-- The difference here is that we add a WHERE clause that filters only results where the column from the right table is NULL. -->

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `LEFT JOIN` (excluding)

- Fetching all the posts that do not have an image attached:

```sql
SELECT p.title, p_i.id
FROM posts as p
LEFT JOIN posts_images as p_i
ON p.id = p_i.post_id
WHERE p_i.id IS NULL;
```

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 28px;
}
code {
  font-size: 70%;
}
</style>

### `LEFT JOIN` (excluding)

- Fetching all the posts that do not have an image attached:

```
                             title                              | id
----------------------------------------------------------------+----
 Fallen Angel : Fact versus Fiction                             |
 Can Blue Bottles Dance : An exploration of Memes               |
 How to Make Your Own Admirable Dress for less than £5          |
 Mickey Mouse - 10 Best Moments                                 |
 From Zero to Shape Shifter - Makeover Tips                     |
 7 Unmissable YouTube Channels About Thoughts                   |
 7 Pictures of Rihanna That We Would Rather Forget              |
 How to Attract More Admirable Subscribers                      |
 10 Awesome Ways to Photograph Blue Bottles                     |
 Introducing programmer - Who Am I And Why Should You Follow Me |
 10 Things Mickey Mouse Can Teach Us About Thoughts             |
 Snakes Are the New Black                                       |
 Blue Bottles Are the New Black                                 |
 The Week: Top Stories About Rihanna                            |
 How to Increase Your Income Using Just Your Knees.             |
 10 Awesome Ways to Photograph Snakes                           |
 7 Pictures of Paul McCartney That We Would Rather Forget       |
 How to Increase Your Income Using Just Your Ankles.            |
 20 Dress Reviews in Tweet Form                                 |
(19 rows)
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `RIGHT JOIN`

- The `RIGHT JOIN` works (as you may have guessed) just like the `LEFT JOIN` but in this case the results returned are from the right table.

- Let's use a Venn diagram to visualize it better...

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `RIGHT JOIN`

![h:400](./assets/right-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `RIGHT JOIN` (excluding)

![h:400](./assets/right-excluding-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `RIGHT JOIN`

```sql
SELECT table_one.column_name, table_two.column_name
FROM table_one
RIGHT JOIN table_two
ON table_one.column_name = table_two.column_name;
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `RIGHT JOIN` (excluding)

```sql
SELECT table_one.column_name, table_two.column_name
FROM table_one
RIGHT JOIN table_two
ON table_one.column_name = table_two.column_name;
WHERE table_one.column_name IS NULL;
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `INNER JOIN`

- Another type of SQL join is the `INNER JOIN` and in this case, the results returned are the ones that have values defined in **both** tables, _i.e._, the intersection between the tables based on a given condition.

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `INNER JOIN`

![h:400](./assets/inner-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `INNER JOIN`

```sql
SELECT table_one.column_name, table_two.column_name
FROM table_one
INNER JOIN table_two
ON table_one.column_name = table_two.column_name;
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `INNER JOIN`

- Let's fetch the users and the cities they live:

```sql
SELECT users.name, cities.name AS city
FROM users
INNER JOIN cities
ON users.city_id = cities.id
LIMIT 15;
```

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 28px;
}
code {
  font-size: 80%;
}
</style>

## `INNER JOIN`

Let's fetch the users and the cities they live:

```sql
   name   |      city
----------+----------------
 Viviana  | Fortaleza
 Callan   | Rio de Janeiro
 Aila     | Fortaleza
 Moses    | Rio de Janeiro
 Amelia   | Brasília
 Chandler | Fortaleza
 Alicia   | Salvador
 Nehemiah | Fortaleza
 Everly   | Salvador
 Kayson   | Brasília
 Imani    | Brasília
 Jamie    | São Paulo
 Ximena   | São Paulo
 Alexis   | Rio de Janeiro
 Estrella | São Paulo
(15 rows)
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `INNER JOIN`

- But is there any city without users associated with?

```sql
SELECT cities.name AS city, users.name
FROM users
RIGHT JOIN cities
ON users.city_id = cities.id
WHERE city IS NULL;
```

<!-- Using a RIGHT JOIN (excluding) to check if there is any city without users associated -->

---

<!-- _class: invert -->

# The `JOIN` statemnet

```sql
blog=# SELECT cities.name AS city, users.name
blog-# FROM users RIGHT JOIN cities
blog-# ON users.city_id = cities.id
blog-# WHERE users.city_id IS NULL;
      city      | name
----------------+------
 Belo Horizonte |
 Manaus         |
 Curitiba       |
 Recife         |
 Goiânia        |
(5 rows)
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

- There is also another type of `JOIN` which fetches **all** the rows from **both** the left and right tables

- This is the `FULL JOIN` and it is basically a `LEFT JOIN`and `RIGHT JOIN` "combined".

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `FULL OUTER JOIN`

![h:400](./assets/full-outer-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `FULL OUTER JOIN` (excluding)

![h:400](./assets/full-excluding-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `FULL OUTER JOIN`

```sql
SELECT table_one.column_name, table_two.column_name
FROM table_one
FULL OUTER JOIN table_two
ON table_one.column_name = table_two.column_name;
```

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `FULL OUTER JOIN` (excluding)

```sql
SELECT table_one.column_name, table_two.column_name
FROM table_one
FULL OUTER JOIN table_two
ON table_one.column_name = table_two.column_name;
WHERE table_one.column_name IS NULL OR table_two.column_name IS NULL;
```

---

# The `JOIN` statement

- There is also another type of join called `CROSS JOIN` which returns the [cartesian product](https://en.wikipedia.org/wiki/Cartesian_product) based on the condition, but it has very specific use cases.

- SQL joins can be "chained", _i.e._, the result of one `JOIN` operation is joined to another table. Let's see one example...

---

<!-- _class: invert -->

# The `JOIN` statement

- In the `blog` database being used for the examples we have a _many-to-many_ relationship between the `posts` and `images` tables that is achieved **through** the `posts_images` _junction_ table.

- What if we want to list the images URL for each of the posts that have at least one image attached?

---

<!-- _class: invert -->

# The `JOIN` statement

- We can achieve that with the following steps:

  1. Do a `INNER JOIN` between the tables `posts` and `posts_images`
  2. Do another `INNER JOIN` between the results and the `images` table.

---

<!-- _class: invert -->

# The `JOIN` statement

- Results goal:

| Post ID | Title      | Images            |
| ------- | ---------- | ----------------- |
| 1       | How to SQL | www.images.com/12 |
| 1       | How to SQL | www.images.com/34 |
| 3       | Git 101    | www.images.com/5  |
| ...     | ...        | ...               |

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

# The `JOIN` statement

### Step 1:

```sql
SELECT posts.id AS "Post ID", posts.title AS "Title"
FROM posts
INNER JOIN posts_images
ON posts.id = posts_images.post_id;
```

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 90%;
}
</style>

#### The `JOIN` statement

```
 Post ID |                         Title
---------+-------------------------------------------------------
      19 | 10 Things Mickey Mouse Can Teach Us About Thoughts
      12 | Mickey Mouse - 10 Best Moments
      24 | 21 Myths About Blue bottles Debunked
      12 | Mickey Mouse - 10 Best Moments
      11 | How to Make Your Own Admirable Dress for less than £5
      21 | 10 Awesome Ways to Photograph Blue Bottles
      27 | Can Blue Bottles Dance : An exploration of Memes
      13 | How to Attract More Admirable Subscribers
      28 | Blue Bottles Are the New Black
      10 | From Zero to Shape Shifter - Makeover Tips
      23 | How to Increase Your Income Using Just Your Knees.
      19 | 10 Things Mickey Mouse Can Teach Us About Thoughts
      27 | Can Blue Bottles Dance : An exploration of Memes
      23 | How to Increase Your Income Using Just Your Knees.
       2 | 7 Pictures of Rihanna That We Would Rather Forget
      13 | How to Attract More Admirable Subscribers
(16 rows)
```

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 150%;
}
</style>

# The `JOIN` statement

### Step 2:

```diff
- SELECT posts.id AS "Post ID", posts.title AS "Title"
+ SELECT posts.id AS "Post ID", posts.title AS "Title", images.img_url as "Images"
FROM posts_images
INNER JOIN posts
ON posts_images.post_id = posts.id
+ INNER JOIN images
+ ON posts_images.image_id = images.id
+ ORDER BY posts.id;
```

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 90%;
}
</style>

#### The `JOIN` statement

```
 Post ID |                             Title                              |          Images
---------+----------------------------------------------------------------+--------------------------
       4 | 21 Myths About Snakes Debunked                                 | https://onlink.site/yQCF
       4 | 21 Myths About Snakes Debunked                                 | https://onlink.site/iVhX
       4 | 21 Myths About Snakes Debunked                                 | https://onlink.site/93iP
       5 | Introducing programmer - Who Am I And Why Should You Follow Me | https://onlink.site/AuUT
       5 | Introducing programmer - Who Am I And Why Should You Follow Me | https://onlink.site/93iP
       7 | Can Snakes Dance : An exploration of Memes                     | https://onlink.site/YMrl
       8 | Snakes Are the New Black                                       | https://onlink.site/AuUT
      12 | Mickey Mouse - 10 Best Moments                                 | https://onlink.site/yQCF
      16 | The Week: Top Stories About Rihanna                            | https://onlink.site/iVhX
      20 | Mistakes That Snakes Make and How to Avoid Them                | https://onlink.site/AuUT
      20 | Mistakes That Snakes Make and How to Avoid Them                | https://onlink.site/93iP
      22 | 7 Pictures of Paul McCartney That We Would Rather Forget       | https://onlink.site/iVhX
      23 | How to Increase Your Income Using Just Your Knees.             | https://onlink.site/93iP
      25 | Introducing database - Who Am I And Why Should You Follow Me   | https://onlink.site/AuUT
      25 | Introducing database - Who Am I And Why Should You Follow Me   | https://onlink.site/YMrl
      26 | Fallen Angel : Fact versus Fiction                             | https://onlink.site/YMrl
(16 rows)
```

<!-- That way we achieve our goal! -->

---

# Querying with SQL

- We have already seen how to filter the resulst and order them (and join tables too!)

- We still have to learn how to generate analytics (at least one way of doing that)

- And that way is using the `GROUP BY` clause...

---

<!-- _class: invert -->

## The `GROUP BY` statement

- We use `GROUP BY` to _group_ rows using an _aggregate function_ on one or more columns.

- Some of the possible aggregate functions are:

  - `COUNT()` - returns the number of rows that meet the condition

  - `MAX()`, `MIN()` and `AVG()` - return the maximum, minimum and average of the values

  - `SUM()` - returns the sum of the values

---

<!-- _class: invert -->

## The `GROUP BY` statement

- With `GROUP BY` we can answer:

  - How many users live in the most inhabited city from the `cities` table?

  - Who is the user with most posts?

---

<!-- _class: invert -->
<style scoped>
  .language-sql {
    font-size: 150%;
  }
</style>

# The `GROUP BY` statement

Basic syntax:

```sql
SELECT AGG(column_name_one), column_name_two
FROM table_name
GROUP BY column_name_two;
```

- Where `AGG()` is an [aggregate function](https://www.postgresql.org/docs/15/functions-aggregate.html)

---

<!-- _class: invert -->
<style scoped>
  .language-sql {
    font-size: 150%;
  }
</style>

# The `GROUP BY` statement

Let's start with a query using a single table:

```sql
SELECT COUNT(image_id), post_id
FROM posts_images
GROUP BY post_id
ORDER BY count DESC;
```

---

<!-- _class: invert -->

<style scoped>
  li {
     font-size: 70%;
  }
  .language-sql {
    font-size: 100%;
  }
</style>

### The `GROUP BY` statement

- The result shows the posts ordered by the number of images attached:

```sql
 count | post_id
-------+---------
     3 |       4
     2 |      25
     2 |      20
     2 |       5
     1 |      12
     1 |      23
     1 |      22
     1 |       8
     1 |      26
     1 |      16
     1 |       7
(11 rows)
```

---

<!-- _class: invert -->
<style scoped>
  .language-sql {
    font-size: 150%;
  }
</style>

# The `GROUP BY` statement

But we can use `GROUP BY` on joins too:

```sql
SELECT COUNT(users.id) AS "Users" , cities.name AS "City"
FROM users
INNER JOIN cities
ON users.city_id = cities.id
GROUP BY cities.name
ORDER BY "Users" DESC;
```

---

<!-- _class: invert -->
<style scoped>
  .language-sql {
    font-size: 150%;
  }
</style>

# The `GROUP BY` statement

And get the number of users per city:

```sql
 Users |      City
-------+----------------
     9 | Salvador
     7 | Brasília
     6 | Fortaleza
     5 | São Paulo
     3 | Rio de Janeiro
(5 rows)
```

---

# Nice Job! :tada:

If you got this far, you already know the core concepts and can do A LOT of things with SQL :sunglasses:

---

# Entering the next level

The next slides will introduce more advanced concepts, all of them are very powerful, and let us make our databases more robust and performant

---

# Entering the next level

- **Subqueries**: using queries within queries.
- **Views**: virtual tables for simplified querying.
- **Indexes**: optimizing data retrieval.
- **Transactions**: ensuring data consistency.
- **Multi-tenancy**: how to separate user environments.

---

<!-- _class: invert -->

# Subqueries

---

<!-- _class: invert -->

# Views

---

<!-- _class: invert -->

# Indexes

---

<!-- _class: invert -->

# Transactions

---

<!-- _class: invert -->

# Multi-tenancy

---

<!-- _class: invert -->

# Agenda

- ![1](https://icongr.am/material/numeric-1-circle.svg?color=ffffff) **~~SQL~~**

- ![2](https://icongr.am/material/numeric-2-circle.svg?color=ffffff) **ORMs**
  - Definition, popular ORMs and examples
- ![3](https://icongr.am/material/numeric-3-circle.svg?color=ffffff) **Best Practices and Tips**
  - Performance, security and debugging

<!-- With the SQL part complete, let's dive in ORMs -->

---

![2 w:128 h:128](https://icongr.am/material/numeric-2-circle.svg?color=666666)

# ORM

**O**bject **R**elational **M**apping

<!-- So what is ORM? It stands for object relational mapping. -->

---

# ![2](https://icongr.am/material/numeric-2-circle.svg?color=666666) ORM

### What is an ORM?

ORM is a technique that uses _object-oriented programming_ to interact with databases.

 <!-- Presenter notes. -->

---

<!-- _class: invert -->

<style scoped>
  table {
    font-size: 80%;
  }
  </style>

##### Here is a list of ORM libraries for different languages

|                                    **Language**                                    |        **ORM Libraries**         |
| :--------------------------------------------------------------------------------: | :------------------------------: |
| ![](https://icongr.am/simple/nodejs.svg?size=45&color=ffffff&colored=false) NodeJS |       Sequelize<br>Prisma        |
| ![](https://icongr.am/simple/python.svg?size=45&color=ffffff&colored=false) Python | SQLAlchemy<br>Django<br>SQLModel |
|     ![](https://icongr.am/simple/go.svg?size=45&color=ffffff&colored=false) Go     |           GORM<br>REL            |
|   ![](https://icongr.am/simple/java.svg?size=45&color=ffffff&colored=false) Java   |     Hibernate<br>EclipseLink     |
|   ![](https://icongr.am/simple/csharp.svg?size=45&color=ffffff&colored=false) C#   |       Entity<br>NHibernate       |
|                                                                                    |                                  |

---

<!-- _class: invert -->

# An Example With Go

Let's build a simple HTTP server that returns a random motivational quote.

We can start with the implementation of the endpoint `/quote` with a hard coded value.

---

<!-- _class: invert -->

Starter code: `main.go`

```go
package main

import (
  "log"
  "net/http"
)

func quoteHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte("Nothing is impossible.\n"))
}

func main() {
  http.HandleFunc("/quote", quoteHandler)

  log.Println("Listening on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

<!-- _class: invert -->

### An Example With Go

- We can spin up the server:

```
$ go run main.go
2023/06/30 00:53:16 Listening on port 8080
```

And make a request:

```
$ curl http://localhost:8080/quote
Nothing is impossible.
```

---

<!-- _class: invert -->

### An Example With Go

- Awesome! We've got our first quote.

- But we will need more quotes to randomize the reponse.

- To do that, instead of a single hardcoded quote, let's create a list of quotes.

---

<!-- _class: invert -->

```go
import (
  "log"
  "math/rand"
  "net/http"
)

var quotes = []string{
  "Nothing is impossible.\n",
  "If you're going through hell, keep going.\n",
  "We need much less than we think we need.\n",
  "If things go wrong, don't go with them.\n",
  "Whatever you are, be a good one.\n",
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
  index := rand.Intn(len(quotes))
  quote := quotes[index]
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte(quote))
}

```

---

<!-- _class: invert -->

### An Example With Go

- Running the server again:

```
$ go run main.go
2023/06/30 00:53:16 Listening on port 8080
```

And making a few requests:

```
$ curl http://localhost:8080/quote
Whatever you are, be a good one.
$ curl http://localhost:8080/quote
If you're going through hell, keep going.
$ curl http://localhost:8080/quote
Nothing is impossible.
```

---

<!-- _class: invert -->

### An Example With Go

- Beautiful! We already are returning random quotes.
- But... this will not scale well, of course. What if we want to store hundreds or even thousands of quotes? :thinking:
- That's when a database comes in handy. In the next step, we are going to connect to a PostgreSQL database so we can separate the concerns.

---

<!-- _class: invert -->

### An Example With Go

<div class="columns red">
<div class="border-right-white">
Current

![h:370](./assets/go-server.png)

</div>
<div>
Goal

![h:370](./assets/go-server-postgres.png)

</div>
</div>

---

<!-- _class: invert -->

### An Example With Go

- In the diagram showed, the Go server and the PostgreSQL server are different services, but will not necessarely run on different computers.

- Before we connect to the database, the database must be ready to receive connections.

- We're going to see how to do it, it is quite simple!

---

<!-- _class: invert -->

### An Example With Go

1. If not installed already, install [PostgreSQL](https://www.postgresql.org/download/)
2. Create a database called `quotes_db` and connect to it:

```bash
psql
DROP DATABASE IF EXISTS quotes_db;
CREATE DATABASE quotes_db;
\c quotes_db
```

---

<!-- _class: invert -->

### An Example With Go

3. Create a table called `quotes` and add some values:

```sql
CREATE TABLE quotes (id serial PRIMARY KEY, quote varchar NOT NULL);
INSERT INTO quotes(quote)
VALUES ('Nothing is impossible'),
    ('If you`re going through hell, keep going'),
    ('We need much less than we think we need'),
    ('If things go wrong, don`t go with them'),
    ('Whatever you are, be a good one');
```

---

<!-- _class: invert -->

### An Example With Go

Here is the data stored ready to be fetched:

```
$ psql quotes_db
quotes_db=# SELECT * FROM quotes;
 id |                  quote
----+------------------------------------------
  1 | Nothing is impossible
  2 | If you`re going through hell, keep going
  3 | We need much less than we think we need
  4 | If things go wrong, don`t go with them
  5 | Whatever you are, be a good one
(5 rows)
```

---

<!-- _class: invert -->

### An Example With Go

- Instead of using `psql` we want to query the database **from the server**.
- To achieve that we need to _connect_ to the database (DB) and then _do the query_ to get a quote e and return it in the response.

---

<!-- _class: invert -->

### An Example With Go

- To connect to the DB we are going to use the `database/sql` native Go package and also a PostgreSQL [driver](https://github.com/golang/go/wiki/SQLDrivers). In this example, [`lib/pq`](https://github.com/lib/pq).

```
go mod init go-orm-example
go get -u github.com/lib/pq
```

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 28px;
}
.language-go {
  font-size:65%;
}
</style>

#### An Example With Go

- Next, let's update the code to create the DB connection:

```go
import (
  "database/sql"
  "log"
  "math/rand"
  "net/http"

  _ "github.com/lib/pq"
)

var db *sql.DB

func init() {
  connStr := "postgresql://localhost/quotes_db?sslmode=disable"
  var err error
  db, err = sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal("Failed to connect to database:", err)
  }

  err = db.Ping()
  if err != nil {
    log.Fatal("Failed to ping database:", err)
  }
  log.Println("Connected to the database")
}
```

---
<!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-go {
  font-size:65%;
}
</style>

#### An Example With Go

- Testing the connection:

```diff
$ go run main.go
+ 2023/07/01 00:20:51 Connected to the database
  2023/07/01 00:20:51 Listening on port 8080
```

- Awesome, we are connected to the DB. Now we can query the table with quotes!

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-go {
  font-size:65%;
}
</style>

#### An Example With Go

- First, we have to define what the query will be. We can either:

  1. Get a list of quotes and randomize the response in the Go server
  2. Get a single random quote from the database

- With option 1 we have a problem: there is a limit on how many quotes we can fetch at once, and it would be really slow and inefficient to ingest several quotes to return just one.
- In option 2 we would have a problem if there was no way to get a random quote from the DB. But as you might have noticed, SQL databases are very flexible and feature-rich. Getting random values is quite easy :tada:

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-sql {
  font-size:150%;
}
</style>

#### An Example With Go

- The query will look like this:

```sql
SELECT quote FROM quotes ORDER BY RANDOM() LIMIT 1;
```

- The `ORDER BY RANDOM()` makes the result random and the `LIMIT 1` clause returns only one result.

- Now we just have to update the Go code and assign this result to one variable to return it in the response...

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 28px;
}
.language-go {
  font-size:90%;
}
</style>

##### An Example With Go

- Add the `getRandomQuote` function to return the result:

```go
func getRandomQuote() (string, error) {
  rows, err := db.Query("SELECT quote FROM quotes ORDER BY RANDOM() LIMIT 1")
  if err != nil {
    return "", err
  }
  defer rows.Close()

  var quote string
  for rows.Next() {
    err := rows.Scan(&quote)
    if err != nil {
      return "", err
    }
  }

  return quote, nil
}
```

---

<!-- _class: invert -->

##### An Example With Go

- Update the `quoteHandler` function:

```go
func quoteHandler(w http.ResponseWriter, r *http.Request) {
  quote, err := getRandomQuote()
  if err != nil {
    log.Println("Failed to retrieve quote:", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte(quote))
}
```

---

<!-- _class: invert -->

##### An Example With Go

- Testing the server again:

```
go run main.go 
2023/07/01 01:28:18 Connected to the database
2023/07/01 01:28:18 Listening on port 8080
```
```
$ curl http://localhost:8080/quote
Nothing is impossible
$ curl http://localhost:8080/quote
We need much less than we think we need
```
:tada: :tada: :tada:

---

### An Example With Go

- So far we have:
  - A Go HTTP server with a single route: `/quote` that returns a random quote.
  - A PostgreSQL server that stores the quotes in a database called `quotes_db`
  - The Go server fetches the quotes from the DB and returns to the client.

**What about the ORM?**

---

### An Example With Go

- Using an ORM is not mandatory. We could improve the example even more and use all the features of SQL using only the current stack.

- But... we can also use an ORM as an abstraction layer between the language we are using in the server (in this example, Go) and SQL.

- That way, we can remove (or at least reduce a lot) the usage of "raw" SQL.

---

<style scoped>
li {
  font-size: 32px;
}
</style>

### PROS and CONS

<div class="columns">
<div class="border-right-black">
<h5>Pros</h5>

- Models are DRY
- SQL injection is harder
- Simpler queries
- Migrations

</div>
<div>
<h5>Cons</h5>

- Complex queries
- Additional tech
- Obfuscates underlying SQL behaviour

</div>
</div>

---

### An Example With Go

- Let's see how we would build the same quotes Go server using an ORM.

- For this example we will be using [GORM](https://gorm.io/)

---

<!-- _class: invert -->

### An Example With Go

- Once again, we start with a simple HTTP server with the `/quote` endpoint:

```go

```

---


<!-- _class: invert -->

# Agenda

- ![1](https://icongr.am/material/numeric-1-circle.svg?color=ffffff) **~~SQL~~**

- ![2](https://icongr.am/material/numeric-2-circle.svg?color=ffffff) **~~ORMs~~**
- ![3](https://icongr.am/material/numeric-3-circle.svg?color=ffffff) **Best Practices and Tips**
  - Performance, security and debugging

<!-- That was a wrap. Let's take a look at the best practices and some tips regarding SQL and ORMs -->

---

![3 w:128 h:128](https://icongr.am/material/numeric-3-circle.svg?color=666666)

# Best Practices and Tips

Things to pay attention!

<!-- So what is ORM? It stands for object relational mapping. -->

---

<style scoped>
section li em {
  font-size: 25px;
}
</style>

## ![3](https://icongr.am/material/numeric-3-circle.svg?color=666666) Best Practices and Tips

- ...

 <!-- Presenter notes. -->
