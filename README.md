# Simple Posts HTTP Server

A lightweight HTTP server written in Go that manages a collection of posts with basic CRUD operations.

## Features

* In-memory storage of posts using a thread-safe map
* RESTful API endpoints for managing posts
* JSON request/response format
* Concurrent request handling with mutex locks

## API Endpoints

### Posts Collection (```/posts```)

#### GET

* Retrieves all posts
* Returns an array of posts in JSON format

#### POST

* Creates a new post
* Request body: JSON object with body field
* Returns the created post with assigned ID
* Status: 201 Created

### Single Post (```/posts/{id}```)

#### GET

* Retrieves a specific post by ID
* Returns the post in JSON format
* Status: 404 if post not found

#### DELETE

* Deletes a specific post by ID
* Status: 200 OK on success
* Status: 404 if post not found

## Data Structure

```go
type Post struct {
    ID   int    `json:id`
    Body string `json:body`
}
```

### Usage

1. Start the server:

```bash
go run main.go
```

2. Example requests:

```bash
# Get all posts
curl http://localhost:8080/posts

# Create a post
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"body": "Hello, World!"}'

# Get specific post
curl http://localhost:8080/posts/1

# Delete post
curl -X DELETE http://localhost:8080/posts/1
```

## Error Handling

* Invalid post ID: 400 Bad Request
* Post not found: 404 Not Found
* Invalid request body: 400 Bad Request
* Method not allowed: 405 Method Not Allowed

## Concurrency

The server uses a mutex to ensure thread-safe access to the posts map, making it safe for concurrent requests.