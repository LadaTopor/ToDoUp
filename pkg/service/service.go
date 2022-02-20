package service

import (
	"github.com/LadaTopor/ToDoUp/pkg/models"
	"github.com/LadaTopor/ToDoUp/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type ToDoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
