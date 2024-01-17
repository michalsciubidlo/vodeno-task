package email

import (
	"context"
	"fmt"
	"strings"
)

type MockService struct{}

func New() *MockService {
	return &MockService{}
}

// SendEmails mocks sending email and prints the payload
func (s *MockService) SendEmails(ctx context.Context, title string, content string, email ...string) error {
	fmt.Printf("sending message to %s:\n[%s]\n%s\n", strings.Join(email, ","), title, content)
	return nil
}
