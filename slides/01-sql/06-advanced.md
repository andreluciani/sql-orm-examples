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

## Subqueries

- _Subqueries_ (also called _inner queries_) are used to query data "inside" another query.

- For instance, we may use subqueries to filter out rows before making a `JOIN` statement, or to define values to be used in thw `WHERE` clause.

- Let's see an example to better understand the concept!

---

<!-- _class: invert -->

## Subqueries

- The basic syntax of subqueries is quite simple, just wrap a regular query in parenthesis:

```sql
SELECT column1, column2
FROM table_name1
WHERE condition
    (SELECT column1, column2
     FROM table_name2);
```

---

<!-- _class: invert -->

## Subqueries

- Here's an example using the same blog dataset used in the "Querying" section

- In this example the query will use a subquery to get all the posts that have authors with names that start with letter "A"

- This query does not look too useful, but it's just an example

---

<!-- _class: invert -->

## Subqueries

- The subquery will look like this:

```sql
SELECT id
FROM users
WHERE name LIKE 'A%';
```

---

<!-- _class: invert -->

## Subqueries

```
 id
----
  3
  5
  7
 14
 17
 21
 24
 26
(8 rows)
```

---

<!-- _class: invert -->

## Subqueries

- Great! Now, we can use the values returned to filter the `posts` table:

```sql
SELECT id,title
FROM posts
WHERE user_id IN
    (SELECT id
    FROM users
    WHERE name LIKE 'A%');
```

---

<!-- _class: invert -->

## Subqueries

```
 id |                            title
----+--------------------------------------------------------------
  9 | 20 Dress Reviews in Tweet Form
 11 | How to Make Your Own Admirable Dress for less than £5
 15 | Unboxing My New Shape Shifter Poo
 18 | 7 Unmissable YouTube Channels About Thoughts
 22 | 7 Pictures of Paul McCartney That We Would Rather Forget
 25 | Introducing database - Who Am I And Why Should You Follow Me
 26 | Fallen Angel : Fact versus Fiction
 28 | Blue Bottles Are the New Black
 29 | 20 Hat Reviews in Tweet Form
 30 | From Zero to Fallen Angel - Makeover Tips
(10 rows)
```

---

<!-- _class: invert -->

## Subqueries

- The previous example was quite simple, but subqueries are very powerful and can be used more than once in a single query

- One important thing to notice is that subqueries are somewhat similar to joins. The same results can be achieved with both features. But a "rule of thumb" is to use subqueries to define _conditions_.

---

<!-- _class: invert -->

# Views

---

<!-- _class: invert -->

## Views

Let's remember one query we've used before to list blog posts titles along with their images

```sql
SELECT posts.id AS "Post ID",
    posts.title AS "Title",
    images.img_url as "Images"
FROM posts_images
    INNER JOIN posts ON posts_images.post_id = posts.id
    INNER JOIN images ON posts_images.image_id = images.id
ORDER BY posts.id;
```

---

<!-- _class: invert -->

## Views

We can create a _view_ to use the data returned exactly as a table. It has a few benefits:

1. Simplifies complex queries
2. Can limit what data is shared with users
3. Provides more meaningful/readable column names
4. Can be used as a step in a really complex query

---

<!-- _class: invert -->

## Views

- When creating views, we must provide a name to it
- View do **not** persist data anywhere, so when querying from a view the underlying query will be executed too

---

<!-- _class: invert -->

## Views

Basic syntax:

```sql
CREATE VIEW view_name AS
-- query to be saved
```

---

<!-- _class: invert -->

## Views

Using the example shown before:

```sql
CREATE VIEW vw_post_titles_images AS
SELECT posts.id AS post_id,
    posts.title AS post_title,
    images.img_url as image_link
FROM posts_images
    INNER JOIN posts ON posts_images.post_id = posts.id
    INNER JOIN images ON posts_images.image_id = images.id
ORDER BY posts.id;
```

---

<!-- _class: invert -->

## Views

After creating the view, we can query as usual:

```
blog=# SELECT * FROM vw_post_titles_images LIMIT 5;
 post_id |                      post_title                       |        image_link
---------+-------------------------------------------------------+--------------------------
       6 | Shape Shifter : Fact versus Fiction                   | https://onlink.site/yQCF
       9 | 20 Dress Reviews in Tweet Form                        | https://onlink.site/93iP
      10 | From Zero to Shape Shifter - Makeover Tips            | https://onlink.site/YMrl
      10 | From Zero to Shape Shifter - Makeover Tips            | https://onlink.site/93iP
      11 | How to Make Your Own Admirable Dress for less than £5 | https://onlink.site/yQCF
(5 rows)
```

---

<!-- _class: invert -->

# Indexes

---

<!-- _class: invert -->

# Transactions

---

<!-- _class: invert -->

# Multi-tenancy
