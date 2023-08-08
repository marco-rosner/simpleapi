# Simple API on Golang

Simple API in Golang for study purpose using Gin Web Framework

## Initializing Server

### go run .

Runs the server on http://localhost:8080

## Endpoints

### JSON routes
    GET - /todos        // get all TODO's
	GET - /todos/:id    // get TODO by ID
	POST - /todos       // post new todo

### Form routes
	GET - /client       // using query string

### URI routes
	GET - /:name/:id    // /:name/:uuid