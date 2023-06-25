# SQL and ORM Examples

Simple repository to illustrate SQL and ORM concepts

The slides are built with [Marp](https://github.com/marp-team/marp-cli)

## **Part 1: SQL**

1. Introduction to SQL:
   - [x]  Definition of SQL (Structured Query Language).
   - [x]  Basic concepts: databases, tables, rows, and columns.
   - [x]  SQL's role in managing relational databases.

2. Data Manipulation Language (DML):
   - [ ]  SELECT statement: retrieving data from tables.
   - [ ]  INSERT statement: inserting data into tables.
   - [ ]  UPDATE statement: modifying existing data.
   - [ ]  DELETE statement: removing data from tables.

3. Data Definition Language (DDL):
   - [ ]  CREATE TABLE: creating a new table.
   - [ ]  ALTER TABLE: modifying table structure.
   - [ ]  DROP TABLE: deleting a table.

4. SQL Relationships:
   - [ ]  ONE-TO-ONE
   - [ ]  ONE-TO-MANY
   - [ ]  MANY-TO-MANY

5. SQL Constraints:
   - [ ]  NOT NULL - Ensures that a column cannot have a NULL value
   - [ ]  UNIQUE - Ensures that all values in a column are different
   - [ ]  PRIMARY KEY - A combination of a NOT NULL and UNIQUE. Uniquely identifies each row in a table
   - [ ]  FOREIGN KEY - Prevents actions that would destroy links between tables
   - [ ]  CHECK - Ensures that the values in a column satisfies a specific condition
   - [ ]  DEFAULT - Sets a default value for a column if no value is specified
   - [ ]  CREATE INDEX - Used to create and retrieve data from the database very quickly

6. Querying with SQL:
   - [ ]  Filtering data with WHERE clause.
   - [ ]  Sorting data with ORDER BY clause.
   - [ ]  Combining data with JOIN clauses.
   - [ ]  Aggregating data with GROUP BY clause.

7. Advanced SQL Concepts:
   - [ ]  Subqueries: using queries within queries.
   - [ ]  Views: virtual tables for simplified querying.
   - [ ]  Indexes: optimizing data retrieval.
   - [ ]  Transactions: ensuring data consistency.
   - [ ]  Multi-tenancy: how to separate user environments.

## **Part 2: Object-Relational Mappers (ORMs)**

1. Introduction to ORMs:
   - [ ]  Definition of ORMs and their purpose.
   - [ ]  Benefits and drawbacks of using ORMs.

2. Popular ORMs:
   - [ ]  Overview of popular ORMs (e.g., Hibernate, SQLAlchemy, Sequelize).
   - [ ]  Language-specific ORMs (e.g., Django ORM, ActiveRecord).

3. Mapping Objects to Tables:
   - [ ]  Object-Relational impedance mismatch.
   - [ ]  Mapping classes/objects to database tables.

4. CRUD Operations with ORMs:
   - [ ]  Creating objects and inserting data.
   - [ ]  Retrieving objects using queries.
   - [ ]  Updating and deleting objects.

5. Relationships and Associations:
   - [ ]  Defining relationships between objects (e.g., one-to-one, one-to-many, many-to-many).
   - [ ]  Handling associations using ORMs.

6. Advanced ORM Concepts:
   - [ ]  Lazy loading vs. eager loading.
   - [ ]  Caching and performance optimization.
   - [ ]  ORM-specific features and configurations.

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

4. Integration with Frameworks:
   - [ ]  Integrating ORMs with popular web frameworks.
   - [ ]  Framework-specific tips and tricks.

5. Troubleshooting and Debugging:
   - [ ]  Common issues with ORMs and SQL.
   - [ ]  Debugging techniques and tools.
