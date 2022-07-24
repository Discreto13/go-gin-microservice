package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/discreto13/go-gin-microservice/internal/service"
	"github.com/discreto13/go-gin-microservice/internal/storage"
	"github.com/discreto13/go-gin-microservice/internal/transport/rest"
	"github.com/discreto13/go-gin-microservice/internal/transport/rest/handler"
)

func main() {
	log.Println("Start server")

	db := preparePostgresDB()
	defer db.Close()

	userStorage := storage.NewUserStorage(db)
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
	os.Exit(0)
}

func preparePostgresDB() *sql.DB {
	configuration := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	log.Println("DB configuration:", configuration)
	db, err := sql.Open("postgres", configuration)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			log.Fatal("failed to connect after 10 seconds")
		}
	}
	log.Println("Connected:", db.Ping() == nil)

	_, err = db.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`CREATE TABLE users (
		id TEXT PRIMARY KEY NOT NULL,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		birthday TEXT NOT NULL
	);`)
	if err != nil {
		panic(err)
	}

	log.Println("Table users is created")

	return db
}
