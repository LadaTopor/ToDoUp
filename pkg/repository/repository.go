package repository

import (
	"github.com/LadaTopor/ToDoUp/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type ToDoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	ToDoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
