package repository

import (
	//"main/model"
	"fmt"
	"main/graph/model"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	GetPosts() ([]*model.Post, error)
	CreatePost(in model.NewPost) (int, error)
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

func (ms *MemoryStorage) GetPosts() ([]*model.Post, error) {
	posts := []*model.Post{}
	for _, p := range ms.posts {
		posts = append(posts, &p)
	}
	return posts, nil
}

func (ms *MemoryStorage) CreatePost(in model.NewPost) (int, error) {
	ms.mu.Lock()
	id := ms.postInc
	ms.postInc++
	ms.posts[ms.postInc] = model.Post{ID: ms.postInc - 1, Title: in.Title + fmt.Sprint(id), Text: in.Text, AuthorID: in.AuthorID, IsCommentsUnabled: in.IsCommentsUnabled}
	ms.mu.Unlock()
	return id, nil
}
