package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/michalsciubidlo/vodeno-task/pkg/customermailing"
)

type service interface {
	Add(ctx context.Context, msg customermailing.MailingMessage) error // Add new MailingMessage
	Send(ctx context.Context, mailingID int) error                     // Send messages to everyone with mailing id
	Delete(ctx context.Context, mailingID int) error                   // Delete all entries older than 5 minutes
}

func SetupRoutes(e *echo.Echo, service service) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/messages", func(c echo.Context) error {
		var m MailingMessage
		if err := c.Bind(&m); err != nil {
			c.Logger().Error("failed to bind request: " + err.Error())
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		msg := toCustomerMailingMessage(m)
		err := service.Add(c.Request().Context(), msg)
		if err != nil {
			c.Logger().Error("failed to add new mailing message: " + err.Error())
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.String(http.StatusCreated, "Created")
	})

	e.DELETE("/api/messages/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "id must be integer")
		}

		err = service.Delete(c.Request().Context(), id)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.String(http.StatusOK, "OK")
	})

	e.POST("/api/messages/send", func(c echo.Context) error {
		var payload MailingIDPayload
		if err := c.Bind(&payload); err != nil {
			c.Logger().Error("failed to bind request: " + err.Error())
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		err := service.Send(c.Request().Context(), payload.MailingID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.String(http.StatusOK, "OK")
	})
}
