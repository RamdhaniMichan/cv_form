## How to run this application

Install the dependencies and start the server.

```sh
git clone 
go mod download
go mod tidy
go run main.go
```

# List Route

Several available lists that can be executed:

| Method | Route | Description |
| ------ | ------ | ------ |
| GET | {{url}}/api/v1/migrate | Execute this route to migrate the database |
| GET | {{url}}/api/v1/books | Get all the book lists in the database |
| POST | {{url}}/api/v1/book | Add book to the database |
| GET | {{url}}/api/v1/book/{id} | Get book details with path parameters |
| GET | {{url}}/api/v1/book/?id= | Get book details with query parameters |
| PUT | {{url}}/api/v1/book/{id} | Change the book details in the database based on id |
| DELETE | {{url}}/api/v1/book/{id} | Delete data in the database based on id |