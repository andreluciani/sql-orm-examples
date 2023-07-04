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
