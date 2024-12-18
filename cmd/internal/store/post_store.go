package store

import (
	"sync"

	"github.com/react-picasso/go-httpserver/internal/models"
)

type PostStore struct {
	posts  map[int]models.Post
	nextID int
	mu     sync.Mutex
}

func NewPostStore() *PostStore {
	return &PostStore{
		posts:  make(map[int]models.Post),
		nextID: 1,
	}
}

func (s *PostStore) List() []models.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	posts := make([]models.Post, 0, len(s.posts))
	for _, p := range s.posts {
		posts = append(posts, p)
	}
	return posts
}

func (s *PostStore) Create(post models.Post) models.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	post.ID = s.nextID
	s.nextID++
	s.posts[post.ID] = post
	return post
}

func (s *PostStore) Get(id int) (models.Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[id]
	return post, exists
}

func (s *PostStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.posts[id]
	if !exists {
		return false
	}

	delete(s.posts, id)
	return true
}
