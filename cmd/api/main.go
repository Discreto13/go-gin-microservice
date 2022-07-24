package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/discreto13/go-gin-microservice/internal/service"
	"github.com/discreto13/go-gin-microservice/internal/storage"
	"github.com/discreto13/go-gin-microservice/internal/transport/rest"
	"github.com/discreto13/go-gin-microservice/internal/transport/rest/handler"
)

func main() {
	log.Println("Start server")

	userStorage := storage.NewUserStorage()
	userService := service.NewUserService(userStorage)
	userHandler := handler.NewUserHandler(userService)

	finalHandler := handler.NewHandler(userHandler)
	restSrv := rest.NewServer(8080, finalHandler.Init())

	go func() {
		if err := restSrv.ListenAndServe(); err != nil {
			log.Printf("serving stopped: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Println("Shutdown request received")

	restSrv.Stop(context.Background())
	// db.Close()
	os.Exit(0)
}
