package main

import (
	"log"

	todoup "github.com/LadaTopor/ToDoUp"
	"github.com/LadaTopor/ToDoUp/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todoup.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
