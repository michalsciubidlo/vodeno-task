package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michalsciubidlo/vodeno-task/pkg/customermailing"
)

//go:generate go run github.com/matryer/moq -out mocks_test.go . service:serviceMock
type service interface {
	Add(ctx context.Context, msg customermailing.MailingMessage) error // Add new MailingMessage
	Send(ctx context.Context, mailingID int) error                     // Send messages to everyone with mailing id
	Delete(ctx context.Context, mailingID int) error                   // Delete all entries with matching mailing id older than 5 minutes
}

type handler struct {
	s service
}

func SetupRoutes(e *echo.Echo, service service) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h := handler{s: service}

	e.POST("/api/messages", h.create)
	e.DELETE("/api/messages/:id", h.delete)
	e.POST("/api/messages/send", h.send)
}
