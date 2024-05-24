package repository

import (
	"main/model"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	GetPosts() ([]model.Post, error)
	CreatePost(in model.Post) error
}

// type Store struct {
// 	db *sql.DB
// }

// func NewStore(db *sql.DB) StoreInterface {
// 	return &Store{
// 		db: db,
// 	}
// }

type MemoryStorage struct {
	mu      sync.RWMutex
	posts   map[int]model.Post
	postInc int
}

func NewMemoryStore() StoreInterface {
	return &MemoryStorage{
		mu:      sync.RWMutex{},
		posts:   map[int]model.Post{},
		postInc: 0,
	}
}

func (ms *MemoryStorage) GetPosts() ([]model.Post, error) {
	posts := []model.Post{}
	for _, p := range ms.posts {
		posts = append(posts, p)
	}
	return posts, nil
}

func (ms *MemoryStorage) CreatePost(in model.Post) error {
	ms.mu.Lock()
	ms.postInc++
	ms.posts[ms.postInc] = in
	ms.mu.Unlock()
	return nil
}
