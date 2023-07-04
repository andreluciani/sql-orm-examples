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
  ![w:1000](/assets/blog-schema.png)
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

![h:400](/assets/left-join.png)

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

![h:400](/assets/left-excluding-join.png)

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

![h:400](/assets/right-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `RIGHT JOIN` (excluding)

![h:400](/assets/right-excluding-join.png)

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

![h:400](/assets/inner-join.png)

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

![h:400](/assets/full-outer-join.png)

---

<!-- _class: invert -->

# The `JOIN` statemnet

## `FULL OUTER JOIN` (excluding)

![h:400](/assets/full-excluding-join.png)

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
