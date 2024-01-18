package email

import (
	"context"
	"strings"

	"github.com/labstack/echo"
)

type MockService struct {
	log echo.Logger
}

func New(log echo.Logger) *MockService {
	return &MockService{
		log: log,
	}
}

// SendEmails mocks sending email and prints the payload
func (s *MockService) SendEmails(ctx context.Context, title string, content string, email ...string) error {
	s.log.Printf("sending message to %s:\n[%s]\n%s\n", strings.Join(email, ","), title, content)
	return nil
}
