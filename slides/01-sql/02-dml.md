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
