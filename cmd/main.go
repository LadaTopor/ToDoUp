package main

import (
	"log"

	todoup "github.com/LadaTopor/ToDoUp"
	"github.com/LadaTopor/ToDoUp/pkg/handler"
	"github.com/LadaTopor/ToDoUp/pkg/repository"
	"github.com/LadaTopor/ToDoUp/pkg/service"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "asdfg",
		DBName:   "tododb",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoup.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
