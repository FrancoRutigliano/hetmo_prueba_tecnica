package main

import (
	"hetmo_prueba_tecnica/config"
	server "hetmo_prueba_tecnica/internal/shared/infrastructure/entrypoint/http_server"
	"log"
)

func main() {
	config, err := config.SetUp()
	if err != nil {
		log.Fatal("error loading .env --> ", err)
	}

	s := server.NewServer(config)
	if err = s.Run(); err != nil {
		log.Fatal("error to inicialize the app --> ", err)
	}
	log.Println("app running succesfully")
}
