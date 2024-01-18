package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) create(c echo.Context) error {
	var m MailingMessage
	if err := c.Bind(&m); err != nil {
		c.Logger().Error("failed to bind request: " + err.Error())
		return c.String(http.StatusBadRequest, "payload is corrupted")
	}

	msg := toCustomerMailingMessage(m)
	err := h.s.Add(c.Request().Context(), msg)
	if err != nil {
		c.Logger().Error("failed to add new mailing message: " + err.Error())
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.String(http.StatusCreated, "Created")
}

func (h *handler) delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "id must be integer")
	}

	err = h.s.Delete(c.Request().Context(), id)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.String(http.StatusOK, "OK")
}

func (h *handler) send(c echo.Context) error {
	var payload MailingIDPayload
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error("failed to bind request: " + err.Error())
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	err := h.s.Send(c.Request().Context(), payload.MailingID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.String(http.StatusOK, "OK")
}
