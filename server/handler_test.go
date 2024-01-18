package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/michalsciubidlo/vodeno-task/pkg/customermailing"
	"github.com/stretchr/testify/require"
)

var (
	messagePayloadJSON = `{
		"email": "bozenka.kowalska@example.com",
		"title": "urgent email",
		"content": "simple text",
		"mailing_id": 1,
		"insert_time": "2021-02-24T01:42:38Z"
	}`
	corruptedMessagePayloadJSON = `{
		"email": "bozenka.kowalska@example.com",
		"title": "urgent email",
		"content": "simple text
		"mailing_id": 1,
		"insert_time": "2021-02-24T01:42:38Z"
	}`
)

func TestAdd(t *testing.T) {
	type testCase struct {
		name           string
		expectedStatus int
		serviceMock    *serviceMock
		request        *http.Request
	}
	testCases := []testCase{
		{
			name:           "should return 201 (Created) if payload is correct",
			expectedStatus: http.StatusCreated,
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(messagePayloadJSON))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			}(),
			serviceMock: &serviceMock{
				AddFunc: func(ctx context.Context, msg customermailing.MailingMessage) error { return nil },
			},
		},
		{
			name:           "should return 400 (Bad Request) if payload is corrupted",
			expectedStatus: http.StatusBadRequest,
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(corruptedMessagePayloadJSON))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			}(),
		},
		{
			name:           "should return 500 (Internal Error) if service returns error",
			expectedStatus: http.StatusInternalServerError,
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(messagePayloadJSON))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			}(),
			serviceMock: &serviceMock{
				AddFunc: func(ctx context.Context, msg customermailing.MailingMessage) error { return fmt.Errorf("test") },
			},
		},
	}

	for _, tc := range testCases {
		rec := httptest.NewRecorder()
		e := echo.New().NewContext(tc.request, rec)
		h := &handler{s: tc.serviceMock}

		require.NoError(t, h.create(e))
		require.Equal(t, tc.expectedStatus, rec.Code)
	}
}
