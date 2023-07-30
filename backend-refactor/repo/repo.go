package repo

import (
	"boilerplate/backend-refactor/model"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type IRepo interface {
	GetUserByName(name string) (model.User, error)
}
