[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/andreluciani/sql-orm-examples)

# SQL and ORM Examples

Simple repository to illustrate SQL and ORM concepts

The slides use PostgreSQL and Go in most of the examples.

The `examples` and `examples-orm` directories have all the code used in the slides.

The slides are built with [Marp](https://github.com/marp-team/marp-cli)

## **Part 1: SQL**

1. Introduction to SQL:
   - [x]  Definition of SQL (Structured Query Language).
   - [x]  Basic concepts: databases, tables, rows, and columns.
   - [x]  SQL's role in managing relational databases.

2. Data Manipulation Language (DML):
   - [x]  SELECT statement: retrieving data from tables.
   - [x]  INSERT statement: inserting data into tables.
   - [x]  UPDATE statement: modifying existing data.
   - [x]  DELETE statement: removing data from tables.

3. Data Definition Language (DDL):
   - [x]  CREATE TABLE: creating a new table.
   - [x]  ALTER TABLE: modifying table structure.
   - [x]  DROP TABLE: deleting a table.

4. SQL Relationships:
   - [x]  ONE-TO-ONE
   - [x]  ONE-TO-MANY
   - [x]  MANY-TO-MANY

5. SQL Constraints:
   - [x]  NOT NULL - Ensures that a column cannot have a NULL value
   - [x]  UNIQUE - Ensures that all values in a column are different
   - [x]  PRIMARY KEY - A combination of a NOT NULL and UNIQUE. Uniquely identifies each row in a table
   - [x]  FOREIGN KEY - Prevents actions that would destroy links between tables
   - [x]  CHECK - Ensures that the values in a column satisfies a specific condition
   - [x]  DEFAULT - Sets a default value for a column if no value is specified
   - [x]  CREATE INDEX - Used to create and retrieve data from the database very quickly

6. Querying with SQL:
   - [x]  Limiting resuts with LIMIT.
   - [x]  Filtering data with WHERE clause.
   - [x]  Sorting data with ORDER BY clause.
   - [x]  Combining data with JOIN clauses.
   - [x]  Aggregating data with GROUP BY clause.

7. Advanced SQL Concepts:
   - [x]  Subqueries: using queries within queries.
   - [x]  Views: virtual tables for simplified querying.
   - [x]  Indexes: optimizing data retrieval.
   - [x]  Transactions: ensuring data consistency.
   - [x]  Multi-tenancy: how to separate user environments.

## **Part 2: Object-Relational Mappers (ORMs)**

1. Introduction to ORMs:
   - [x]  Definition of ORMs and their purpose.
   - [x]  Benefits and drawbacks of using ORMs.
   - [x]  Mapping classes/objects to database tables.

2. CRUD Operations with ORMs:
   - [x]  Creating objects and inserting data.
   - [x]  Retrieving objects using queries.
   - [x]  Updating and deleting objects.

4. Additional ORM Concepts:
   - [ ]  Lazy loading vs. eager loading.
   - [ ]  Caching and performance optimization.
   - [ ]  Object-Relational impedance mismatch

## **Part 3: Best Practices and Tips**

1. Efficient Querying:
   - [ ]  Writing efficient SQL queries.
   - [ ]  Using indexes and query optimization techniques.

2. Security Considerations:
   - [ ]  Preventing SQL injection attacks.
   - [ ]  Data sanitization and validation.

3. Performance Tuning:
   - [ ]  Optimizing database schema and indexes.
   - [ ]  Caching strategies and query optimization.

5. Troubleshooting and Debugging:
   - [ ]  Common issues with ORMs and SQL.
   - [ ]  Debugging techniques and tools.
