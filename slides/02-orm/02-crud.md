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
}
```