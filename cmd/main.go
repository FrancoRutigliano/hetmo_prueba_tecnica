package main

import (
	"hetmo_prueba_tecnica/config"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/data"
	server "hetmo_prueba_tecnica/internal/shared/infrastructure/entrypoint/http_server"
	"log"
)

func main() {
	config, err := config.SetUp()
	if err != nil {
		log.Fatal("error loading .env --> ", err)
	}

	db, err := data.GetConnection(config)
	if err != nil {
		log.Fatal("error connecting db --> ", err)
	}

	log.Println("db connected succesfully")

	s := server.NewServer(config, db)
	if err = s.Run(); err != nil {
		log.Fatal("error to inicialize the app --> ", err)
	}
	log.Println("app running succesfully")
}
