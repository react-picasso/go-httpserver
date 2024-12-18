package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/react-picasso/go-httpserver/internal/models"
	"github.com/react-picasso/go-httpserver/internal/store"
)

var postStore = store.NewPostStore()

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetPosts(w, r)
	case http.MethodPost:
		handleCreatePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SinglePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGetPost(w, r, id)
	case http.MethodDelete:
		handleDeletePost(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts := postStore.List()
	sendJSON(w, posts, http.StatusOK)
}

func handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &post); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	post = postStore.Create(post)
	sendJSON(w, post, http.StatusCreated)
}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {
	post, exists := postStore.Get(id)
	if !exists {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	sendJSON(w, post, http.StatusOK)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {
	if ok := postStore.Delete(id); !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
