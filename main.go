package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michalsciubidlo/vodeno-task/pkg/customermailing"
	"github.com/michalsciubidlo/vodeno-task/pkg/email"
	"github.com/michalsciubidlo/vodeno-task/server"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DB connection
	db, err := sqlx.Connect("postgres", "postgres://postgres:example@postgres_db:5432/api?sslmode=disable")
	if err != nil {
		log.Fatal("failed to create db connection: " + err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping DB" + err.Error())
	}

	// Initialize service
	emailService := email.New(e.Logger)
	service := customermailing.NewService(emailService, customermailing.NewStorage(db))

	// Setup api routes
	server.SetupRoutes(e, service)

	// Start server
	e.Logger.Print("hello swiecie!")
	e.Logger.Fatal(e.Start(":8080"))
}
