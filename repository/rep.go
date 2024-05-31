package repository

import (
	//"main/model"
	"fmt"
	"main/graph/model"
	"sync"

	Errors "main/errors"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	GetPosts() ([]*model.Post, error)
	CreatePost(in model.NewPost) (int, error)
	GetPost(id int) (*model.Post, error)
	SwitchCommentsCreation(id int) error
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
	posts := make([]*model.Post, len(ms.posts))
	idx := 0
	for _, val := range ms.posts {
		posts[idx] = &model.Post{ID: val.ID, Title: val.Title, Text: val.Text, AuthorID: val.AuthorID, IsCommentsUnabled: val.IsCommentsUnabled}
		idx++
	}
	//log.Println(posts)
	return posts, nil
}

func (ms *MemoryStorage) CreatePost(in model.NewPost) (int, error) {
	ms.mu.Lock()
	id := ms.postInc
	ms.posts[ms.postInc] = model.Post{ID: ms.postInc, Title: in.Title + fmt.Sprint(id), Text: in.Text, AuthorID: in.AuthorID, IsCommentsUnabled: in.IsCommentsUnabled}
	ms.postInc++
	ms.mu.Unlock()
	return id, nil
}

func (ms *MemoryStorage) GetPost(id int) (*model.Post, error) {
	if res, ok := ms.posts[id]; ok {
		return &res, nil
	}
	return nil, Errors.ErrNotFound404
}

func (ms *MemoryStorage) SwitchCommentsCreation(id int) error {
	ms.mu.Lock()
	ms.posts[id] = model.Post{ID: ms.posts[id].ID, Title: ms.posts[id].Title, Text: ms.posts[id].Text, AuthorID: ms.posts[id].AuthorID,
		IsCommentsUnabled: !ms.posts[id].IsCommentsUnabled}
	ms.mu.Unlock()
	return nil
}
