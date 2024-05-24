package usecase

import (
	"main/model"

	rep "main/repository"
)

type UsecaseInterface interface {
	GetPosts() ([]model.Post, error)
	CreatePost(in model.Post) error
}

type Usecase struct {
	store rep.StoreInterface
}

func NewUsecase(s rep.StoreInterface) UsecaseInterface {
	return &Usecase{
		store: s,
	}
}

func (uc *Usecase) GetPosts() ([]model.Post, error) {
	return uc.store.GetPosts()
}
func (uc *Usecase) CreatePost(in model.Post) error {
	return uc.store.CreatePost(in)
}
