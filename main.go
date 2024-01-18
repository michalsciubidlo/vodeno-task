package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
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
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:example@postgres_db:5432/api")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(context.Background()); err != nil {
		e.Logger.Fatal("failed to ping DB: " + err.Error())
	}

	rows, err := conn.Query(context.Background(), "SELECT datname FROM pg_database;")
	if err != nil {
		e.Logger.Fatal("failed to query DB: " + err.Error())
	}
	for rows.Next() {
		var str string
		if err := rows.Scan(&str); err != nil {
			e.Logger.Fatal("failed to scan row: " + err.Error())
		}
		e.Logger.Print("database: " + str)
	}

	// Initialize service
	emailService := email.New()
	service := customermailing.NewService(emailService, customermailing.NewStorage())

	server.SetupRoutes(e, service)

	// Start server
	e.Logger.Print("hello swiecie!")
	e.Logger.Fatal(e.Start(":8080"))
}
