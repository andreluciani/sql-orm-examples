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

## Views

- When creating views, the `view_name` must be unique in the database
- A good approach is to prefix views names with `vw_`

---

<!-- _class: invert -->

# Indexes

---

<!-- _class: invert -->

## Indexes

- In SQL databases indexes are used to query data with **better performance**.

- They work similarly to an index of a book: instead of scanning through _all_ the pages of the book searching for a keyword,
  we can look at the index and find what we're looking for way faster.

---

<!-- _class: invert -->

## Indexes

- Indexes may be used for columns that are queried a lot.
- Indexes are created automatically for _primary keys_ or _unique_ constraint columns (it makes sense, since primary keys or unique values are often used in SQL queries)

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

## Indexes

- Basic syntax:

```sql
CREATE INDEX index_name
ON table_name (column1, column2, ...);
```

And just like that, you've created an index!

---

<!-- _class: invert -->

## Indexes

- When creating indexes, the `index_name` must be unique in the database
- A good approach is to prefix indexes names with `idx_`
- There are different _types_ of indexes, and they vary among SQL DMS, but the default is [B-tree](https://en.wikipedia.org/wiki/B-tree)

---

<!-- _class: invert -->
<style scoped>
.language-sql {
  font-size: 120%;
}
</style>

## Indexes

- Let's create an index for the `blog`database:

```sql
CREATE INDEX idx_posts_titles
ON posts(title);
```

In this example, querying for keywords in the `title` column would be faster because of the new index.

---

# Why not create indexes on **every** column?

---

#### Why not create indexes on every column?

1. Indexes store copies of data
2. If the table is not big, it may be faster to scan the table than to scan the index
3. Operations such as `INSERT`, `UPDATE` and `DELETE` take longer when indexes are used

---

<!-- _class: invert -->

## Indexes

"Rules of thumb" for index creation:

- Columns used in `WHERE`, `JOIN` or `HAVING` clauses
- Foreign keys columns
- Columns with several distinct values

---

<!-- _class: invert -->

# Transactions

---

<!-- _class: invert -->

## Transactions

- _Transactions_ allow grouping multiple database operations in a "all-or-nothing" manner

- It is a fundamental concept that enhances integrity and consistency of the database

---

<!-- _class: invert -->

## Transactions

- For instance when a user creates a profile in our `blog` database, it might be the case that the city where they live is not listed in the `cities` table and has to be created in the same operation.

- Let's see what the query would look like...

---

<!-- _class: invert -->

## Transactions

```sql
INSERT INTO cities(name)
VALUES ('New York');
INSERT INTO users(name, email, city_id, role_id)
VALUES ('John', 'john@email.com', (SELECT city_id FROM cities WHERE name = 'New York'), 1);
```

---

## Transactions

- What if the query fails at the creation of the `cities` entry?

- And what about if it fails at the user creation?

- In both cases, inconsistencies could arise...

---

<!-- _class: invert -->

## Transactions

That's why transactions are so important, if we wrap the operations in a single transaction, the database will persist the data only if **both** operations are successful!

---

<!-- _class: invert -->

## Transactions

- Here's the basic syntax:

```sql
BEGIN;
-- SQL statements
COMMIT;
```

- Or:

```sql
BEGIN;
-- SQL statements
ROLLBACK;
```

---

<!-- _class: invert -->

## Transactions

- When using `COMMIT` the operations will be persisted to the database.

- `ROLLBACK` on the other hand, discards the changes up to the point where the transacion begun (at the `BEGIN` statement)

- It is possible to `ROLLBACK` to different steps of the transaction using _savepoints_.

---

<!-- _class: invert -->

## Transactions

```sql
BEGIN;
UPDATE accounts SET balance = balance - 100.00
    WHERE name = 'Alice';
SAVEPOINT my_savepoint;
UPDATE accounts SET balance = balance + 100.00
    WHERE name = 'Bob';
-- oops ... forget that and use Wally's account
ROLLBACK TO my_savepoint;
UPDATE accounts SET balance = balance + 100.00
    WHERE name = 'Wally';
COMMIT;
```

---

<!-- _class: invert -->

## Transactions

- Here's our example using a transaction

```sql
BEGIN;
INSERT INTO cities(name)
VALUES ('New York');
INSERT INTO users(name, email, city_id, role_id)
VALUES ('John', 'john@email.com', (SELECT city_id FROM cities WHERE name = 'New York'), 1);
COMMIT;
```

---

<!-- _class: invert -->

# Multi-tenancy

---

<!-- _class: invert -->

## Multi-tenancy

-  When developing a software that will be used by several users, it might be the case that we need to _separate concerns_, i.e., customer data, based on a context.

- For instance, if the software is a B2B SaaS, we probably want to separate data from each client (company).

- To achieve this, there are a few options that will be discussed next. When this scenario shows up, we call each customer context a _tenant_

---

<!-- _class: invert -->

## Multi-tenant Architectures

Here are a few options to achieve multi-tenancy:

1. Database per tenant
2. Shared tables
    2.1. Tenant identification column
    2.2. Row-level access
3. Schema per tenant

---

<!-- _class: invert -->

## Database per Tenant

In the database per tenant architecture, each tenant is allocated a dedicated database instance, ensuring strict data isolation and autonomy. This approach offers excellent security and scalability, making it ideal for large-scale applications with diverse tenants. However, it can be resource-intensive and costly to maintain numerous separate databases.

---

<!-- _class: invert -->

## Shared Tables

Shared tables employ a single database with a common schema, where a unique identifier (usually a tenant ID) distinguishes tenant-specific data. Rows are tagged with this identifier, enabling efficient data separation. While this approach reduces infrastructure complexity and costs, it demands careful management to prevent unauthorized data access and can face performance challenges as the dataset grows.

---

<!-- _class: invert -->

## Schema per Tenant

In the schema per tenant strategy, every tenant gets their distinct schema within a shared database. This method combines aspects of both previous approaches, granting tenants autonomy over their data structure while benefiting from resource consolidation. Nonetheless, it requires careful schema management and may still incur additional overhead compared to shared tables.
