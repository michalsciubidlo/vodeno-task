package customermailing

import "context"

type emailService interface {
	SendEmails(ctx context.Context, title string, content string, email ...string) error
}

type storage interface {
	// Add(ctx context.Context, item Message)
}

type Service struct {
	emailService emailService
	storage      storage
}

func NewService(emailService emailService, storage storage)

func (s *Service) Add()
