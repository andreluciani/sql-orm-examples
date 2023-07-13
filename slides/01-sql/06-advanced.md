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

* _Subqueries_ (also called _inner queries_) are used to query data "inside" another query.

* For instance, we may use subqueries to filter out rows before making a `JOIN` statement, or to define values to be used in thw `WHERE` clause.

* Let's see an example to better understand the concept!


---

<!-- _class: invert -->

## Subqueries

* The basic syntax of subqueries is quite simple, just wrap a regular query in parenthesis:

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

* Here's an example using the same blog dataset used in the  "Querying" section

* In this example the query will use a subquery to get all the posts that have authors with names that start with letter "A"

* This query does not look too useful, but it's just an example



---

<!-- _class: invert -->

## Subqueries

* The subquery will look like this:

```sql
SELECT id
FROM users
WHERE name LIKE 'A%';
```

---

<!-- _class: invert -->

## Subqueries

* Great! Now, we can use the values returned to filter the `posts` table:

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

* The previous example was quite simple, but subqueries are very powerful and can be used more than once in a single query

* One important thing to notice is that subqueries are somewhat similar to joins. The same results can be achieved with both features. But a "rule of thumb" is to use subqueries to define _conditions_.
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
