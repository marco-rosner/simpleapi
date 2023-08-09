# Simple API on Golang

Simple API in Golang for study purpose using Gin Web Framework

## Server

### Manual initialization

```go
go run .
```

Runs the server on `http://localhost:8080`

### Live reload

Install Gin

```go
go install github.com/codegangsta/gin@latest
```

Make sure that Gin is in the `GOPATH/bin` and run:

```go
gin run main.go
```

For every change that you made in your code, Gin will rebuild your application and start a proxy on `http://localhost:3000` that needs to be requested to start the application server. In other words, you will need to send a request to `http://localhost:3000` then another request to the application server (`http://localhost:8080`) to test your changes.

## Endpoints

### JSON routes

    GET - /todos        // get all TODO's
    GET - /todos/:id    // get TODO by ID
    POST - /todos       // post new todo

### Form routes

    GET - /client       // using query string

### URI routes

    GET - /:name/:id    // /:name/:uuid
