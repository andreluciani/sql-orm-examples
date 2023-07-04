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

![h:550](/assets/example-database.png)

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

![h:190](/assets/tables.png)

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
