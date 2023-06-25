---
marp: true
title: SQL and ORM
description: Illustrating fundamentals of SQL and ORM
theme: base
transition: fade
paginate: false
_paginate: false
style: |
  .columns {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 1rem;
  }
math: mathjax
---

<!-- class: lead -->

# <!--fit--> Fundamentals of SQL and ORM

Understanding SQL at its Core: Building a Strong Foundation for Data Management

<!-- Presenter notes can be written HTML comments. -->

---

<!-- class: lead gaia -->

# <!--fit--> SQL

---

<!-- class: lead -->

##### <!--fit--> What is SQL?

##### <!--fit--> SQL stands for Structured Query Language<br />and is the most common language used in databases.

---

<!-- transition: cover -->

Changed the kind of transition to `cover`.

---

<!-- _transition: none -->

Disabled transition for this slide.

~~Hi~~ Hello, ~there~ world!

---

<!-- class: gaia -->
<!-- _transition: clockwise -->

### This is what clockwise looks like!

<div class="columns">
<div>

Javascrpt:

```js
console.log("Hello, JavaScript!");
```

Python:

```python
print("Hello, Python!")
```

Go:

```go
fmt.Println("Hello, Go!")
```

</div>
<div>

Bash:

```bash
psql
CREATE my_database;
q\
```

SQL:

```sql
SELECT * FROM users WHERE id=5;
```

</div>
</div>

---

<!-- class: invert -->
<!-- Regular list -->

- One
- Two
- Three

1. One
2. Two
3. Three

---

<!-- class: none -->
<!-- Fragmented list -->

- One
- Two
- Three

1. One
2. Two
3. Three

---

# Today's topics

- ![1](https://icongr.am/material/numeric-1-circle.svg?color=666666) Introduction
- ![2](https://icongr.am/material/numeric-2-circle.svg?color=666666) Features
- ![3](https://icongr.am/material/numeric-3-circle.svg?color=666666) Conclusion

---

<!-- _class: lead -->

![1 w:256 h:256](https://icongr.am/material/numeric-1-circle.svg?color=ff9900)

# Introduction

---

# ![1](https://icongr.am/material/numeric-1-circle.svg?color=666666) Introduction

Marp is an open-sourced Markdown presentation ecosystem.

It provides a writing experience of presentation slides by Markdown.

---

# Title with $x=2$ equation

Render inline math such as $ax^2+bc+c$.

$$ I\_{xx}=\int\int_Ry^2f(x,y)\cdot{}dydx $$

$$
f(x) = \int_{-\infty}^\infty
    \hat f(\xi)\,e^{2 \pi i \xi x}
    \,d\xi
$$

---

# PSQL - PostgreSQL CLI

```bash
psql
CREATE DATABASE blog;
CREATE DATABASE supermarket;
CREATE DATABASE gym;
\l
```

```bash
                                 List of databases
    Name     | Owner  | Encoding |   Collate   |    Ctype    | Access privileges
-------------+--------+----------+-------------+-------------+-------------------
 blog        | gitpod | UTF8     | en_US.UTF-8 | en_US.UTF-8 |
 gym         | gitpod | UTF8     | en_US.UTF-8 | en_US.UTF-8 |
 supermarket | gitpod | UTF8     | en_US.UTF-8 | en_US.UTF-8 |

```

---

<!-- _class: invert -->

### Databases, Tables, Rows and Columns

#### List tables of a database

```bash
psql
\c blog
\dt
```

```
          List of relations
 Schema |    Name    | Type  | Owner
--------+------------+-------+--------
 public | posts      | table | gitpod
 public | user_post  | table | gitpod
 public | user_types | table | gitpod
 public | users      | table | gitpod

```

 <!-- Each table has columns, which are attributes of every entry (row) of the table. You can think of tables as spreadsheets. Each column has a type (int, char, datetime, enum, etc), and can have a constraint as well (unique, not null, etc). If the column is a **primary** key, it means that the column identifies the row of the table. If the column is a **foreign** key, that means it is related to ANOTHER table, and it uses the primary key OR another unique column of the other table. -->

---

# SQL Constraints

NOT NULL - Ensures that a column cannot have a NULL value
UNIQUE - Ensures that all values in a column are different
PRIMARY KEY - A combination of a NOT NULL and UNIQUE. Uniquely identifies each row in a table
FOREIGN KEY - Prevents actions that would destroy links between tables
CHECK - Ensures that the values in a column satisfies a specific condition
DEFAULT - Sets a default value for a column if no value is specified
CREATE INDEX - Used to create and retrieve data from the database very quickly
