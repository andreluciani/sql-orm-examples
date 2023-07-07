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

# Agenda

- ![1](https://icongr.am/material/numeric-1-circle.svg?color=ffffff) **~~SQL~~**

- ![2](https://icongr.am/material/numeric-2-circle.svg?color=ffffff) **ORMs**
  - Definition, popular ORMs and examples
- ![3](https://icongr.am/material/numeric-3-circle.svg?color=ffffff) **Best Practices and Tips**
  - Performance, security and debugging

<!-- With the SQL part complete, let's dive in ORMs -->

---

![2 w:128 h:128](https://icongr.am/material/numeric-2-circle.svg?color=666666)

# ORM

**O**bject **R**elational **M**apping

<!-- So what is ORM? It stands for object relational mapping. -->

---

# ![2](https://icongr.am/material/numeric-2-circle.svg?color=666666) ORM

### What is an ORM?

ORM is a technique that uses _object-oriented programming_ to interact with databases.

 <!-- Presenter notes. -->

---

<!-- _class: invert -->

<style scoped>
  table {
    font-size: 80%;
  }
  </style>

##### Here is a list of ORM libraries for different languages

|                                    **Language**                                    |        **ORM Libraries**         |
| :--------------------------------------------------------------------------------: | :------------------------------: |
| ![](https://icongr.am/simple/nodejs.svg?size=45&color=ffffff&colored=false) NodeJS |       Sequelize<br>Prisma        |
| ![](https://icongr.am/simple/python.svg?size=45&color=ffffff&colored=false) Python | SQLAlchemy<br>Django<br>SQLModel |
|     ![](https://icongr.am/simple/go.svg?size=45&color=ffffff&colored=false) Go     |           GORM<br>REL            |
|   ![](https://icongr.am/simple/java.svg?size=45&color=ffffff&colored=false) Java   |     Hibernate<br>EclipseLink     |
|   ![](https://icongr.am/simple/csharp.svg?size=45&color=ffffff&colored=false) C#   |       Entity<br>NHibernate       |
|                                                                                    |                                  |

---

<!-- _class: invert -->

# An Example With Go

Let's build a simple HTTP server that returns a random motivational quote.

We can start with the implementation of the endpoint `/quote` with a hard coded value.

---

<!-- _class: invert -->

Starter code: `main.go`

```go
package main

import (
  "log"
  "net/http"
)

func quoteHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte("Nothing is impossible.\n"))
}

func main() {
  http.HandleFunc("/quote", quoteHandler)

  log.Println("Listening on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

<!-- _class: invert -->

### An Example With Go

- We can spin up the server:

```
$ go run main.go
2023/06/30 00:53:16 Listening on port 8080
```

And make a request:

```
$ curl http://localhost:8080/quote
Nothing is impossible.
```

---

<!-- _class: invert -->

### An Example With Go

- Awesome! We've got our first quote.

- But we will need more quotes to randomize the reponse.

- To do that, instead of a single hardcoded quote, let's create a list of quotes.

---

<!-- _class: invert -->

```go
import (
  "log"
  "math/rand"
  "net/http"
)

var quotes = []string{
  "Nothing is impossible.\n",
  "If you're going through hell, keep going.\n",
  "We need much less than we think we need.\n",
  "If things go wrong, don't go with them.\n",
  "Whatever you are, be a good one.\n",
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
  index := rand.Intn(len(quotes))
  quote := quotes[index]
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte(quote))
}

```

---

<!-- _class: invert -->

### An Example With Go

- Running the server again:

```
$ go run main.go
2023/06/30 00:53:16 Listening on port 8080
```

And making a few requests:

```
$ curl http://localhost:8080/quote
Whatever you are, be a good one.
$ curl http://localhost:8080/quote
If you're going through hell, keep going.
$ curl http://localhost:8080/quote
Nothing is impossible.
```

---

<!-- _class: invert -->

### An Example With Go

- Beautiful! We already are returning random quotes.
- But... this will not scale well, of course. What if we want to store hundreds or even thousands of quotes? :thinking:
- That's when a database comes in handy. In the next step, we are going to connect to a PostgreSQL database so we can separate the concerns.

---

<!-- _class: invert -->

### An Example With Go

<div class="columns red">
<div class="border-right-white">
Current

![h:370](/assets/go-server.png)

</div>
<div>
Goal

![h:370](/assets/go-server-postgres.png)

</div>
</div>

---

<!-- _class: invert -->

### An Example With Go

- In the diagram showed, the Go server and the PostgreSQL server are different services, but will not necessarely run on different computers.

- Before we connect to the database, the database must be ready to receive connections.

- We're going to see how to do it, it is quite simple!

---

<!-- _class: invert -->

### An Example With Go

1. If not installed already, install [PostgreSQL](https://www.postgresql.org/download/)
2. Create a database called `quotes_db` and connect to it:

```bash
psql
DROP DATABASE IF EXISTS quotes_db;
CREATE DATABASE quotes_db;
\c quotes_db
```

---

<!-- _class: invert -->

### An Example With Go

3. Create a table called `quotes` and add some values:

```sql
CREATE TABLE quotes (id serial PRIMARY KEY, quote varchar NOT NULL);
INSERT INTO quotes(quote)
VALUES ('Nothing is impossible'),
    ('If you`re going through hell, keep going'),
    ('We need much less than we think we need'),
    ('If things go wrong, don`t go with them'),
    ('Whatever you are, be a good one');
```

---

<!-- _class: invert -->

### An Example With Go

Here is the data stored ready to be fetched:

```
$ psql quotes_db
quotes_db=# SELECT * FROM quotes;
 id |                  quote
----+------------------------------------------
  1 | Nothing is impossible
  2 | If you`re going through hell, keep going
  3 | We need much less than we think we need
  4 | If things go wrong, don`t go with them
  5 | Whatever you are, be a good one
(5 rows)
```

---

<!-- _class: invert -->

### An Example With Go

- Instead of using `psql` we want to query the database **from the server**.
- To achieve that we need to _connect_ to the database (DB) and then _do the query_ to get a quote and return it in the response.

---

<!-- _class: invert -->

### An Example With Go

- To connect to the DB we are going to use the `database/sql` native Go package and also a PostgreSQL [driver](https://github.com/golang/go/wiki/SQLDrivers). In this example, [`lib/pq`](https://github.com/lib/pq).

```
go mod init go-sql-example
go get -u github.com/lib/pq
```

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 28px;
}
.language-go {
  font-size:65%;
}
</style>

#### An Example With Go

- Next, let's update the code to create the DB connection:

```go
import (
  "database/sql"
  "log"
  "math/rand"
  "net/http"

  _ "github.com/lib/pq"
)

var db *sql.DB

func init() {
  connStr := "postgresql://localhost/quotes_db?sslmode=disable"
  var err error
  db, err = sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal("Failed to connect to database:", err)
  }

  err = db.Ping()
  if err != nil {
    log.Fatal("Failed to ping database:", err)
  }
  log.Println("Connected to the database")
}
```

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-go {
  font-size:65%;
}
</style>

#### An Example With Go

- Testing the connection:

```bash
$ go run main.go
 2023/07/01 00:20:51 Connected to the database
 2023/07/01 00:20:51 Listening on port 8080
```

- Awesome, we are connected to the DB. Now we can query the table with quotes!

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-go {
  font-size:65%;
}
</style>

#### An Example With Go

- First, we have to define what the query will be. We can either:

  1. Get a list of quotes and randomize the response in the Go server
  2. Get a single random quote from the database

- With option 1 we have a problem: there is a limit on how many quotes we can fetch at once, and it would be really slow and inefficient to ingest several quotes to return just one.
- In option 2 we would have a problem if there was no way to get a random quote from the DB. But as you might have noticed, SQL databases are very flexible and feature-rich. Getting random values is quite easy :tada:

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-sql {
  font-size:150%;
}
</style>

#### An Example With Go

- The query will look like this:

```sql
SELECT quote FROM quotes ORDER BY RANDOM() LIMIT 1;
```

- The `ORDER BY RANDOM()` makes the result random and the `LIMIT 1` clause returns only one result.

- Now we just have to update the Go code and assign this result to one variable to return it in the response...

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 28px;
}
.language-go {
  font-size:90%;
}
</style>

##### An Example With Go

- Add the `getRandomQuote` function to return the result:

```go
func getRandomQuote() (string, error) {
  rows, err := db.Query("SELECT quote FROM quotes ORDER BY RANDOM() LIMIT 1")
  if err != nil {
    return "", err
  }
  defer rows.Close()

  var quote string
  for rows.Next() {
    err := rows.Scan(&quote)
    if err != nil {
      return "", err
    }
  }

  return quote, nil
}
```

---

<!-- _class: invert -->

##### An Example With Go

- Update the `quoteHandler` function:

```go
func quoteHandler(w http.ResponseWriter, r *http.Request) {
  quote, err := getRandomQuote()
  if err != nil {
    log.Println("Failed to retrieve quote:", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte(quote))
}
```

---

<!-- _class: invert -->

##### An Example With Go

- Testing the server again:

```
go run main.go
2023/07/01 01:28:18 Connected to the database
2023/07/01 01:28:18 Listening on port 8080
```

```
$ curl http://localhost:8080/quote
Nothing is impossible
$ curl http://localhost:8080/quote
We need much less than we think we need
```

:tada: :tada: :tada:

---

### An Example With Go

- So far we have:
  - A Go HTTP server with a single route: `/quote` that returns a random quote.
  - A PostgreSQL server that stores the quotes in a database called `quotes_db`
  - The Go server fetches the quotes from the DB and returns to the client.

**What about the ORM?**

---

### An Example With Go

- Using an ORM is not mandatory. We could improve the example even more and use all the features of SQL using only the current stack.

- But... we can also use an ORM as an abstraction layer between the language we are using in the server (in this example, Go) and SQL.

- That way, we can remove (or at least reduce a lot) the usage of "raw" SQL.

---

<style scoped>
li {
  font-size: 32px;
}
</style>

### PROS and CONS

<div class="columns">
<div class="border-right-black">
<h5>Pros</h5>

- Models are DRY
- SQL injection is harder
- Simpler queries
- Migrations

</div>
<div>
<h5>Cons</h5>

- Complex queries
- Additional tech
- Obfuscates underlying SQL behaviour

</div>
</div>

---

### An Example With Go

- Let's see how we would build the same quotes Go server using an ORM.

- For this example we will be using [GORM](https://gorm.io/)

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 32px;
}
.language-go {
  font-size: 70%;
}
</style>

### An Example With Go

- Once again, we start with a simple HTTP server with the `/quote` endpoint:

```go
package main

import (
  "log"
  "net/http"
)

func quoteHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte("Nothing is impossible.\n"))
}

func main() {
  http.HandleFunc("/quote", quoteHandler)

  log.Println("Listening on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

<!-- _class: invert -->
<style scoped>
li {
  font-size: 28px;
}

code {
  font-size: 150%;
}
</style>

### An Example With Go

- To connect to the DB we are going to use the `GORM` library with its PostgreSQL driver, [`gorm.io/driver/postgres`](https://github.com/go-gorm/postgres).

```
go mod init go-orm-example
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

---

<!-- _class: invert -->

### An Example With Go

- In the example with Go and vanilla SQL we created the database `quotes_db` and added some data using `psql`.

- Here we are going to use `GORM` to do this step as well, so we can see how the DDL is translated to the ORM.

---

<!-- _class: invert -->

```go
// package, import ...
type Quotes struct {
  ID    uint `gorm:"primaryKey"`
  Quote string
}

var (
  initialQuotes = []Quotes{
    {Quote: "Nothing is impossible"},
    {Quote: "If you`re going through hell, keep going"},
    {Quote: "We need much less than we think we need"},
    {Quote: "If things go wrong, don`t go with them"},
    {Quote: "Whatever you are, be a good one"},
  }
)

func main() {
  dsn := "host=localhost dbname=quotes_db port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&Quotes{})

  for _, quote := range initialQuotes {
    db.Create(&quote)
  }
}
```

<!-- Here's what the seed file looks like. -->

---

<!-- _class: invert -->
<style scoped>
  img {
    padding: 0;
    margin: 0;
  }
  .bottom-space {
  padding-top: 28px;
  padding-bottom: 28px;
}
</style>

<div class="columns">
<div>

<img src="https://icongr.am/simple/go.svg?size=100&color=ffffff&colored=false" />

```go
// package, import ...
type Quotes struct {
  ID    uint `gorm:"primaryKey"`
  Quote string
}

var (
  initialQuotes = []Quotes{
    {Quote: "Nothing is impossible"},
    {Quote: "If you`re going through hell, keep going"},
    {Quote: "We need much less than we think we need"},
    {Quote: "If things go wrong, don`t go with them"},
    {Quote: "Whatever you are, be a good one"},
  }
)

func main() {
  dsn := "host=localhost dbname=quotes_db port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&Quotes{})

  for _, quote := range initialQuotes {
    db.Create(&quote)
  }
}
```

</div>
<div>

<div class="bottom-space">
SQL
</div>

```sql
CREATE TABLE quotes (id serial PRIMARY KEY, quote varchar NOT NULL);
INSERT INTO quotes(quote)
VALUES ('Nothing is impossible'),
    ('If you`re going through hell, keep going'),
    ('We need much less than we think we need'),
    ('If things go wrong, don`t go with them'),
    ('Whatever you are, be a good one');
```

</div>
</div>

<!-- Comparing to raw SQL (not including the creation of the database). As you can see, using Go, the code is more verbose. -->
<!-- At the same time, though, theres's nothing SQL in the code, it is "pure" Go. It uses structs, methods and Go types. -->

---

### An Example With Go

Now, let's update the server to query the database using `GORM`!

---

<!-- _class: invert -->
<style scoped>
  li {
    font-size: 28px;
  }
</style>

### An Example With Go

- First, we will declare a global variable `db` that holds the database connection, and declare the model used in this example:

```go
var db *gorm.DB

type Quotes struct {
  ID    uint `gorm:"primaryKey"`
  Quote string
}
```

- In order to make the code more DRY we could have used a different file to declare the model and use it both for the seed script and the server itself

---

<!-- _class: invert -->
<style scoped>
  .language-go {
    font-size: 150%;
  }
</style>

### An Example With Go

- Then we create a function to execute the database query:

```go
func getRandomQuote() (string, error) {
  var quote Quotes
  err := db.Order("RANDOM()").Take(&quote).Error
  if err != nil {
    return "", err
  }
  return quote.Quote, nil
}
```

---

<!-- _class: invert -->

### An Example With Go

- In the previous slide we have used the `db` variable (which has the `gorm.DB` type) to make the query.

- Two methods were called: `Order()` and `Take()`. The first one is directly related to the SQL statement `ORDER BY`.

- In the [`GORM` docs](https://gorm.io/docs/) we can find how to create models, associations, queries, and so on.

---

<!-- _class: invert -->
<style scoped>
  li {
    font-size: 28px;
  }
</style>

### An Example With Go

- Here are a few "equivalent" functions:

|    **GORM**    |   **SQL**    |
| :------------: | :----------: |
|   `Create()`   |   `INSERT`   |
|   `Where()`    |   `WHERE`    |
|   `Select()`   |   `SELECT`   |
|   `Order()`    |  `ORDER BY`  |
|   `Joins()`    | `LEFT JOIN`  |
| `InnerJoins()` | `INNER JOIN` |

---

<!-- _class: invert -->
<style scoped>
  li {
    font-size: 28px;
  }
</style>

### An Example With Go

- This does **not** mean that the ORM queries will look exactly like the SQL queries. Some SQL keywords are added to the query "behind the scenes". Here's another `GORM` example:

`db.First(&user)` is equivalent to

`SELECT * FROM users ORDER BY id LIMIT 1;`

---

<!-- _class: invert -->
<style scoped>
  li {
    font-size: 28px;
  }
</style>

### An Example With Go

- Actually, `GORM` provides a method to get the resulting SQL query:

```go
sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
  return tx.Order("RANDOM()").Take(&quote)
})
log.Println(sql)
```

Returns:

`SELECT * FROM "quotes" ORDER BY RANDOM() LIMIT 1`

- Quite similar to the query used in the Go + SQL example!
