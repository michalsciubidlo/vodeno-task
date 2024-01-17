package main

import (
	"database/sql"
	"log"

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
	connStr := "host=db user=postgres password=example sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		e.Logger.Fatal("failed to ping DB: " + err.Error())
	}

	// Initialize service
	emailService := email.New()
	service := customermailing.NewService(emailService, customermailing.NewStorage())

	server.SetupRoutes(e, service)

	// Start server
	e.Logger.Print("hello swiecie!")
	e.Logger.Fatal(e.Start(":8080"))
}
