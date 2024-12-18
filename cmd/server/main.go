package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/react-picasso/go-httpserver/internal/handlers"
)

func main() {
	http.HandleFunc("/posts", handlers.PostsHandler)
	http.HandleFunc("/posts/", handlers.SinglePostHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
