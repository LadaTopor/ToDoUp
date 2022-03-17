package main

import (
	"log"
	"os"

	todoup "github.com/LadaTopor/ToDoUp"
	"github.com/LadaTopor/ToDoUp/pkg/handler"
	"github.com/LadaTopor/ToDoUp/pkg/repository"
	"github.com/LadaTopor/ToDoUp/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

	if err := initConfig(); err != nil {
		log.Panic(err)
	} 

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoup.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
