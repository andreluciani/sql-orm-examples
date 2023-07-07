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

### CRUD with `GORM`

- Now, let's take another step and implement a GO API with `GORM` to do CRUD operations. Here's the schema to be implemented:

![h:300](/assets/books-schema.png)

---

<!-- _class: invert -->

### CRUD with `GORM`

- It is a really simple schema to store books and authors with a single `one-to-many` relation (one author can have several books associated with them)

- The first step is to create the models using `GORM`

---

<!-- _class: invert -->

#### CRUD with `GORM`

- `models/authors.go`

```go
package model

import (
  "gorm.io/gorm"
)

type Author struct {
  gorm.Model
  FirstName string
  LastName  string
  Books     []Book
}
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

- `models/books.go`

```go
package model

import (
  "gorm.io/gorm"
)

type Book struct {
  gorm.Model
  Title             string
  Description       string
  YearOfPublication int
  AuthorID          uint
  Author            Author
}
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

- In the previous slides, we have used the [`gorm.Model`](https://gorm.io/docs/models.html#gorm-Model) struct to abstract a lot of things.

- No need to explicitly declare the `id`, `created_at`, `updated_at` and `deleted_at` fields! :tada:

---

#### CRUD with ~~`GORM`~~ [Prisma](https://www.prisma.io/)

- Just for comparison, here is what the exact same models look like using Prisma (a TypeScript ORM)

```javascript
model Authors {
    id        Int       @id @default(autoincrement())
    firstName String?
    lastName  String?
    createdAt DateTime  @default(now())
    updatedAt DateTime  @updatedAt
    deletedAt DateTime?
    books     Books[]
}
```

---

#### CRUD with ~~`GORM`~~ [Prisma](https://www.prisma.io/)

- Just for comparison, here is what the exact same models look like using Prisma (a TypeScript ORM)

```javascript
model Books {
    id                Int       @id @default(autoincrement())
    title             String
    description       String?
    yearOfPublication Int
    author            Authors   @relation(fields: [authorId], references: [id])
    authorId          Int
    createdAt         DateTime  @default(now())
    updatedAt         DateTime  @updatedAt
    deletedAt         DateTime?
}
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

- Back to `GORM` ...

- Let's create a seed file to add some initial data as well. We could have started with empty tables, too, but adding will make simpler to explain the next steps.

---

<!-- _class: invert -->

#### CRUD with `GORM`

```go
package main

import (
  "log"

  "go-book-server/model"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var (
  initialAuthors = []model.Author{
    {FirstName: "William", LastName: "Shakespeare"},
    {FirstName: "Harper", LastName: "Lee"},
  }

  initialBooks = []model.Book{
    {
      Title:             "Macbeth",
      Description:       "A Scottish general's ruthless quest for power...",
      YearOfPublication: 1600,
      AuthorID:          1,
    },
    {
      Title:             "Romeo and Juliet",
      Description:       " The forbidden love between two young individuals...",
      YearOfPublication: 1595,
      AuthorID:          1,
    },
    {
      Title:             "To Kill a Mockingbird",
      Description:       "Set in the racially-charged 1930s Deep South...",
      YearOfPublication: 1860,
      AuthorID:          2,
    },
  }
)
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

```go
func main() {
  dsn := "host=localhost port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  // Checking if DB exists
  rs := db.Raw("SELECT * FROM pg_database WHERE datname = 'books_db';")
  if rs.Error != nil {
    log.Fatal("Raw query failed:", err)
  }

  // If not, create it
  var rec = make(map[string]interface{})
  if rs.Find(rec); len(rec) == 0 {
    if rs := db.Exec("CREATE DATABASE books_db;"); rs.Error != nil {
      log.Fatal("Couldn't create database: ", err)
    }

    // Close db connection
    sql, err := db.DB()
    defer func() {
      _ = sql.Close()
    }()
    if err != nil {
      log.Fatal("An error occurred: ", err)
    }
  }

  // Reconnect and add initial data
  dsn = "host=localhost dbname=books_db port=5432 sslmode=disable"
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&model.Author{}, &model.Book{})

  for _, author := range initialAuthors {
    db.Create(&author)
  }
  for _, book := range initialBooks {
    db.Create(&book)
  }

  log.Println("Successfully added seed data!")
}
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

- Awesome! Now we have a database called `books_db` with a few entries to work with.

```
$ go run seed/main.go
2023/07/05 01:05:05 Successfully added seed data!
$ psql books_db
books_db=# SELECT * FROM authors;
 id |          created_at           |          updated_at           | deleted_at | first_name |  last_name
----+-------------------------------+-------------------------------+------------+------------+-------------
  1 | 2023-07-05 01:05:05.803875+00 | 2023-07-05 01:05:05.803875+00 |            | William    | Shakespeare
  2 | 2023-07-05 01:05:05.805375+00 | 2023-07-05 01:05:05.805375+00 |            | Harper     | Lee
(2 rows)
books_db=# SELECT id,title FROM books;
 id |         title
----+-----------------------
  1 | Macbeth
  2 | Romeo and Juliet
  3 | To Kill a Mockingbird
(3 rows)
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

- Next step: implement CRUD operations.
  - *C*reate
  - *R*etrieve
  - *U*pdate
  - *D*elete

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- We'll start with the _Retrieve_ operation, implementing the `/authors` and `/books` `GET` endpoints:

| **HTTP Method** |  **Endpoint**   |      **Description**      |
| :-------------: | :-------------: | :-----------------------: |
|       GET       |   `/authors`    | Returns a list of authors |
|       GET       | `/authors/<id>` | Returns a specific author |
|       GET       |    `/books`     |  Returns a list of books  |
|       GET       |  `/books/<id>`  |  Returns a specific book  |

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- The code will be structered as follows:

```
.
└── books/
    ├── handler/
    │   ├── authors.go
    │   ├── books.go
    │   └── handler.go
    ├── model/
    │   ├── authors.go
    │   └── books.go
    ├── seed/
    │   └── main.go
    ├── go.mod
    ├── go.sum
    └── main.go
```

- The `model` and `seed` folders are done already

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

```go
package main

import (
  "go-book-server/handler"
  "log"
  "net/http"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var db *gorm.DB

func main() {
  dsn := "host=localhost dbname=books_db port=5432 sslmode=disable"
  var err error
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect to database")
  }

  controller := handler.NewController(db)

  http.HandleFunc("/authors", controller.Authors())
  http.HandleFunc("/authors/", controller.AuthorsByID())

  log.Println("Server started on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

`handler/handler.go`

```go
package handler

import "gorm.io/gorm"

func NewController(db *gorm.DB) *Controller {
  return &Controller{
    db: *db,
  }
}

type Controller struct {
  db gorm.DB
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

`handler/authors.go`

```go
package handler

import (
  "encoding/json"
  "go-book-server/model"
  "log"
  "net/http"
)

func (c *Controller) Authors() http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
      c.ListAuthors(w, r)
      return
    }
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("Method not allowed"))
  })
}
// ...
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

`handler/authors.go`

```go
// ...
func (c *Controller) AuthorsByID() http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
      c.GetAuthorByID(w, r)
      return
    }
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("Method not allowed"))
  })
}
// ...
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

`handler/authors.go`

```go
// ...
func (c *Controller) ListAuthors(w http.ResponseWriter, r *http.Request) {
  var authors []model.Author
  err := c.db.Preload("Books").Find(&authors).Error
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  result, err := json.Marshal(authors)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(result)
}
// ...
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

`handler/authors.go`

```go
// ...
func (c *Controller) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Path[len("/authors/"):]
  var author model.Author
  err := c.db.Preload("Books").First(&author, id).Error
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  result, err := json.Marshal(author)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(result)
}
```

---

<!-- _class: invert -->
<style scoped>
  code {
    font-size: 80%;
  }
  </style>

### CRUD with `GORM`

- In the previous slides, the method `.Preload("Books")` was called.

- This is a feature from `GORM` that tells the query to return the author associated with the book.

- This technique is called _eager loading_ and will be exaplained in more detail later

---

<!-- _class: invert -->

#### CRUD with `GORM`

- Now let's spin up the server and do some testing!

```bash
$ go run main.go
2023/07/05 02:46:57 Server started on http://localhost:8080
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

- `/authors` endpoint:

```bash
$ curl http://localhost:8080/authors | jq
[
   {
    "ID": 1,
    "CreatedAt": "2023-07-06T01:18:28.220539Z",
    "UpdatedAt": "2023-07-06T01:18:28.220539Z",
    "DeletedAt": null,
    "FirstName": "William",
    "LastName": "Shakespeare",
    "Books": [{...},{...}]
  },
  {
    "ID": 2,
    "CreatedAt": "2023-07-06T01:18:28.221913Z",
    "UpdatedAt": "2023-07-06T01:18:28.221913Z",
    "DeletedAt": null,
    "FirstName": "Harper",
    "LastName": "Lee",
    "Books": [{...}]
  }
]
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 75%;
}
</style>

#### CRUD with `GORM`

- `/authors/<id>` endpoint:

```bash
$ curl http://localhost:8080/authors/2 | jq
{
  "ID": 2,
  "CreatedAt": "2023-07-06T01:18:28.221913Z",
  "UpdatedAt": "2023-07-06T01:18:28.221913Z",
  "DeletedAt": null,
  "FirstName": "Harper",
  "LastName": "Lee",
  "Books": [
    {
      "ID": 3,
      "CreatedAt": "2023-07-06T01:18:28.226073Z",
      "UpdatedAt": "2023-07-06T01:18:28.226073Z",
      "DeletedAt": null,
      "Title": "To Kill a Mockingbird",
      "Description": "Set in the racially-charged 1930s ...",
      "YearOfPublication": 1860,
      "AuthorID": 2,
    }
  ]
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

- `/authors/<id>` endpoint (wrong id):

```bash
$ curl http://localhost:8080/authors/123
curl: (52) Empty reply from server
```

- Meanwhile, on the server:

```bash
2023/07/05 02:51:45 /book/handler/authors.go:53 record not found
[3.437ms] [rows:0] SELECT * FROM "authors" WHERE "authors"."id" = '123'
                   AND "authors"."deleted_at" IS NULL
                   ORDER BY "authors"."id" LIMIT 1
2023/07/05 02:51:45 record not found
exit status 1
```

- We can fix this!

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

`handler/authors.go`

```diff
var author model.Author
err := c.db.First(&author, id).Error
if err != nil {
+  if errors.Is(err, gorm.ErrRecordNotFound) {
+    w.WriteHeader(http.StatusNotFound)
+    w.Write([]byte("author not found."))
+    return
+  }
  w.WriteHeader(http.StatusInternalServerError)
  log.Fatal(err)
  return
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

- The `/books` endpoint is implemented in a very similar fashion:

```go
// package, import, var ...

func main() {
  // db connection ...

  controller := handler.NewController(db)

  http.HandleFunc("/authors", controller.Authors())
  http.HandleFunc("/authors/", controller.AuthorsByID())

  http.HandleFunc("/books", controller.Books())
  http.HandleFunc("/books/", controller.BooksByID())

  log.Println("Server started on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

#### CRUD with `GORM`

- The `/books` endpoint is implemented in a very similar fashion:

```go
// package, import, var ...

func main() {
  // db connection ...

  controller := handler.NewController(db)

  http.HandleFunc("/authors", controller.Authors())
  http.HandleFunc("/authors/", controller.AuthorsByID())

  http.HandleFunc("/books", controller.Books())
  http.HandleFunc("/books/", controller.BooksByID())

  log.Println("Server started on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

<!-- _class: invert -->

##### CRUD with `GORM`

```go
// ...
func (c *Controller) ListBooks(w http.ResponseWriter, r *http.Request) {
  var Books []model.Book
  err := c.db.Preload("Author").Find(&Books).Error
  if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      w.WriteHeader(http.StatusNotFound)
      w.Write([]byte("Book not found."))
      return
    }
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  result, err := json.Marshal(Books)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(result)
}
```

---

<!-- _class: invert -->

##### CRUD with `GORM`

```go
// ...
func (c *Controller) GetBookByID(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Path[len("/Books/"):]
  var Book model.Book
  err := c.db.Preload("Author").First(&Book, id).Error
  if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      w.WriteHeader(http.StatusNotFound)
      w.Write([]byte("Book not found."))
      return
    }
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  result, err := json.Marshal(Book)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(result)
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

- Testing the `/books` endpoint:

<div class="columns">
<div>

```bash
$ curl http:/localhost:8080/books | jq
[
  {
    "ID": 1,
    "CreatedAt": "2023-07-05T03:15:43.252328Z",
    "UpdatedAt": "2023-07-05T03:15:43.252328Z",
    "DeletedAt": null,
    "Title": "Macbeth",
    "Description": "A Scottish general's ruthless quest ...",
    "YearOfPublication": 1600,
    "AuthorID": 1,
    "Author": {
      "ID": 1,
      "CreatedAt": "2023-07-05T03:15:43.248912Z",
      "UpdatedAt": "2023-07-05T03:15:43.248912Z",
      "DeletedAt": null,
      "FirstName": "William",
      "LastName": "Shakespeare"
    }
  },
  {
    "ID": 2,
    "CreatedAt": "2023-07-05T03:15:43.254232Z",
    "UpdatedAt": "2023-07-05T03:15:43.254232Z",
    "DeletedAt": null,
    "Title": "Romeo and Juliet",
    "Description": " The forbidden love between two young individuals ...",
    "YearOfPublication": 1595,
    "AuthorID": 1,
    "Author": {
      "ID": 1,
      "CreatedAt": "2023-07-05T03:15:43.248912Z",
      "UpdatedAt": "2023-07-05T03:15:43.248912Z",
      "DeletedAt": null,
      "FirstName": "William",
      "LastName": "Shakespeare"
    }
  },
  ...
```

</div>
<div>

```bash
...
  {
    "ID": 3,
    "CreatedAt": "2023-07-05T03:15:43.255502Z",
    "UpdatedAt": "2023-07-05T03:15:43.255502Z",
    "DeletedAt": null,
    "Title": "To Kill a Mockingbird",
    "Description": "Set in the racially-charged 1930s Deep South...",
    "YearOfPublication": 1860,
    "AuthorID": 2,
    "Author": {
      "ID": 2,
      "CreatedAt": "2023-07-05T03:15:43.250817Z",
      "UpdatedAt": "2023-07-05T03:15:43.250817Z",
      "DeletedAt": null,
      "FirstName": "Harper",
      "LastName": "Lee"
    }
  }
]
```

</div>
</div>

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 80%;
}
</style>

##### CRUD with `GORM`

- Testing the `/books/<id>` endpoint:

```bash
$ curl http:/localhost:8080/books/1 | jq
{
  "ID": 1,
  "CreatedAt": "2023-07-05T03:15:43.252328Z",
  "UpdatedAt": "2023-07-05T03:15:43.252328Z",
  "DeletedAt": null,
  "Title": "Macbeth",
  "Description": "A Scottish general's ruthless quest...",
  "YearOfPublication": 1600,
  "AuthorID": 1,
  "Author": {
    "ID": 1,
    "CreatedAt": "2023-07-05T03:15:43.248912Z",
    "UpdatedAt": "2023-07-05T03:15:43.248912Z",
    "DeletedAt": null,
    "FirstName": "William",
    "LastName": "Shakespeare"
  }
}
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

- That concludes the _Retrieve_ operation. Next: _Delete_
  - *C*reate
  - ~~*R*etrieve~~ :ballot_box_with_check:
  - *U*pdate
  - *D*elete

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- For the _Delete_ operation, we will implement a [soft delete](https://gorm.io/docs/delete.html#Soft-Delete) for the endpoints the `/authors/<id>` and `/books/<id>`:

| **HTTP Method** |  **Endpoint**   |      **Description**      |
| :-------------: | :-------------: | :-----------------------: |
|     DELETE      | `/authors/<id>` | Deletes a specific author |
|     DELETE      |  `/books/<id>`  |  Deletes a specific book  |

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

```diff
func (c *Controller) AuthorsByID() http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
      c.GetAuthorByID(w, r)
      return
    }
+    if r.Method == http.MethodDelete {
+      c.DeleteAuthor(w, r)
+      return
+    }
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("Method not allowed"))
  })
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

```go
func (c *Controller) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Path[len("/authors/"):]
  var author = model.Author{}
  err := c.db.Where("id = ?", id).Delete(&author).Error
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Testing the new endpoint:

```bash
$ curl -i -X DELETE http://localhost:8080/authors/1
HTTP/1.1 204 No Content
```

- Awesome! For the `/books/<id>` endpoint, the same logic can be used, changing only the model used in the `GORM` query.

---

<!-- _class: invert -->

#### CRUD with `GORM`

- All right, one more operation done! Next: _Create_
  - *C*reate
  - ~~*R*etrieve~~ :ballot_box_with_check:
  - *U*pdate
  - ~~*D*elete~~ :ballot_box_with_check:

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- For the _Create_ operation, we will use the `POST` HTTP method in the endpoints `/authors` and `/books`:

| **HTTP Method** | **Endpoint** |  **Description**  |
| :-------------: | :----------: | :---------------: |
|      POST       |  `/authors`  | Inserts an author |
|      POST       |   `/books`   |  Inserts a book   |

- Also, when requesting these endpoints, we'll need a payload with the data to be added.

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

```diff
func (c *Controller) Authors() http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
+    if r.Method == http.MethodPost {
+      c.CreateAuthor(w, r)
+      return
+    }
    if r.Method == http.MethodGet {
      c.ListAuthors(w, r)
      return
    }
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("Method not allowed"))
  })
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

```go
func (c *Controller) CreateAuthor(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  var payload createAuthorPayload
  if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
  }
  author := &model.Author{
    FirstName: payload.FirstName,
    LastName:  payload.LastName,
  }
  if err := c.db.Create(&author).Error; err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
  }
  result, err := json.Marshal(author)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  w.WriteHeader(http.StatusCreated)
  w.Write(result)
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- The `createAuthorPayload` type is declared with the fields and types used in the payload:

```go
type createAuthorPayload struct {
  FirstName string
  LastName  string
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Let's test the route:

```bash
$ curl -X POST 'http://localhost:8080/authors' \
  -H 'Content-Type: application/json' \
  -d '{"firstName":"Jane","lastName":"Austen"}' | jq
{
  "ID": 9,
  "CreatedAt": "2023-07-06T23:44:41.681442697Z",
  "UpdatedAt": "2023-07-06T23:44:41.681442697Z",
  "DeletedAt": null,
  "FirstName": "Jane",
  "LastName": "Austen",
  "Books": null
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Let's test the route:

```bash
$ curl -X POST 'http://localhost:8080/authors' \
  -H 'Content-Type: application/json' \
  -d '{"firstName":123,"lastName":"Austen"}'
json: cannot unmarshal number into Go struct field createAuthorPayload.FirstName of type string
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

- The implementation for `/books` is quite similar. Here's the `createBookPayload` type declaration:

```go
type createBookPayload struct {
	Title             string
	Description       string
	YearOfPublication int
	AuthorID          int
}
```

---

<!-- _class: invert -->

#### CRUD with `GORM`

- One more operation to go! Next: _Update_
  - ~~*C*reate~~ :ballot_box_with_check:
  - ~~*R*etrieve~~ :ballot_box_with_check:
  - *U*pdate
  - ~~*D*elete~~ :ballot_box_with_check:

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- For the _Update_ operation, we will use the `PATCH` HTTP method in the endpoints `/authors` and `/books`:

| **HTTP Method** | **Endpoint** |  **Description**  |
| :-------------: | :----------: | :---------------: |
|     PATCH       |  `/authors`  | Updates an author |
|     PATCH       |   `/books`   |  Updates a book   |

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

```diff
func (c *Controller) AuthorsByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			c.GetAuthorByID(w, r)
			return
		}
+		if r.Method == http.MethodPatch {
+			c.UpdateAuthor(w, r)
+			return
+		}
		if r.Method == http.MethodDelete {
			c.DeleteAuthor(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	})
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- First we get the author being updated and check if it exists:

```go
func (c *Controller) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/authors/"):]
	var author model.Author
	err := c.db.Preload("Books").First(&author, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("author not found."))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
// ...
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Then we load the payload and update the fields:

```go
// ...
	defer r.Body.Close()
	var payload updateAuthorPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if payload.FirstName != "" {
		author.FirstName = payload.FirstName
	}

	if payload.LastName != "" {
		author.LastName = payload.LastName
	}

	if err := c.db.Save(&author).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
// ...
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Lastly we return the author with updated fields:

```go
// ...
	result, err := json.Marshal(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- The `updateAuthorPayload` type is identical with the `createAuthorPayload` type:

```go
type updateAuthorPayload struct {
  FirstName string
  LastName  string
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Testing the route:

```bash
$ curl -X PATCH http://localhost:8080/authors/1 \
  -H 'Content-Type: application/json' \
  -d '{"firstName":"Gulielmus","lastName":"Shakspere"}' | jq
{
  "ID": 1,
  "CreatedAt": "2023-07-06T01:18:28.220539Z",
  "UpdatedAt": "2023-07-07T00:55:05.965733711Z",
  "DeletedAt": null,
  "FirstName": "Gulielmus",
  "LastName": "Shakspere",
  "Books": [...]
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Testing the route:

```bash
$ curl -X PATCH http://localhost:8080/authors/1 \
  -H 'Content-Type: application/json' \
  -d '{"firstName":"William"}' | jq
{
  "ID": 1,
  "CreatedAt": "2023-07-06T01:18:28.220539Z",
  "UpdatedAt": "2023-07-07T00:55:05.965733711Z",
  "DeletedAt": null,
  "FirstName": "William",
  "LastName": "Shakspere",
  "Books": [...]
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Testing the route:

```bash
$ curl -X PATCH http://localhost:8080/authors/1 \
  -H 'Content-Type: application/json' \
  -d '{"lastName":"Shakespeare"}' | jq
{
  "ID": 1,
  "CreatedAt": "2023-07-06T01:18:28.220539Z",
  "UpdatedAt": "2023-07-07T00:55:05.965733711Z",
  "DeletedAt": null,
  "FirstName": "William",
  "LastName": "Shakespeare",
  "Books": [...]
}
```

---

<!-- _class: invert -->
<style scoped>
li,code,td,th {
  font-size: 90%;
}
</style>

#### CRUD with `GORM`

- Let's test the route:

```bash
$ curl -X PATCH http://localhost:8080/authors/1 \
  -H 'Content-Type: application/json' \
  -d '{"firstName":123}'
json: cannot unmarshal number into Go struct field updateAuthorPayload.FirstName of type string
```
---

<!-- _class: invert -->

#### CRUD with `GORM`

- That completes the CRUD operations :tada:
  - ~~*C*reate~~ :ballot_box_with_check:
  - ~~*R*etrieve~~ :ballot_box_with_check:
  - ~~*U*pdate~~ :ballot_box_with_check:
  - ~~*D*elete~~ :ballot_box_with_check:

