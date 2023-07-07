# Get
curl http://localhost:8080/authors/1 | jq
curl http://localhost:8080/authors | jq
curl http://localhost:8080/books/1 | jq
curl http://localhost:8080/books | jq
# Post
curl -X POST http://localhost:8080/authors -H 'Content-Type: application/json'  -d '{"firstName":"Jane","lastName":"Austen"}' | jq
curl -X POST http://localhost:8080/books -H 'Content-Type: application/json'  -d '{"title":"Hamlet","description":"Prince Hamlet seeks revenge for his fathers murder, wrestling with his own sanity and the complexities of life, culminating in a tragic finale","yearOfPublication": 1603, "authorID": 1}' | jq
# Patch
curl -X PATCH http://localhost:8080/authors/1 -H 'Content-Type: application/json'  -d '{"firstName":"Gulielmus","lastName":"Shakspere"}' | jq
curl -X PATCH http://localhost:8080/authors/1 -H 'Content-Type: application/json'  -d '{"firstName":"William"}' | jq
curl -X PATCH http://localhost:8080/authors/1 -H 'Content-Type: application/json'  -d '{"lastName":"Shakespeare"}' | jq
curl -X PATCH http://localhost:8080/authors/1 -H 'Content-Type: application/json'  -d '{"firstName":123}' | jq
# Delete
curl -i -X DELETE http://localhost:8080/authors/10