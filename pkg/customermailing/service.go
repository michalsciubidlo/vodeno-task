package customermailing

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type emailService interface {
	SendEmails(ctx context.Context, title string, content string, email ...string) error
}

type storage interface {
	Add(ctx context.Context, item MailingMessage) error
	DeleteOlderThan(ctx context.Context, t time.Time) error
	GetMailingMessagesByID(ctx context.Context, id int) ([]MailingMessage, error)
}

type Service struct {
	emailService emailService
	storage      storage
}

func NewService(emailService emailService, storage storage) *Service {
	return &Service{
		emailService: emailService,
		storage:      storage,
	}
}

// Add saves a new MailingMessage record
func (s *Service) Add(ctx context.Context, item MailingMessage) error {
	item.InsertTime = time.Now()
	err := s.storage.Add(ctx, item)
	if err != nil {
		return errors.Wrap(err, "failed to add mailing message")
	}
	return nil
}

// Delete removes all MailingMessages older than 5 minutes
func (s *Service) Delete(ctx context.Context) error {
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
	err := s.storage.DeleteOlderThan(ctx, fiveMinutesAgo)
	if err != nil {
		return errors.Wrap(err, "failed to delete messages")
	}
	return nil
}

// Send sends messages to all with mailing id
func (s *Service) Send(ctx context.Context, mailingID int) error {
	msgs, err := s.storage.GetMailingMessagesByID(ctx, mailingID)
	if err != nil {
		return errors.Wrap(err, "failed to get messages by mailing id")
	}

	// todo: safe fall back when sending fails
	for _, msg := range msgs {
		err := s.emailService.SendEmails(ctx, msg.Title, msg.Content, msg.Email)
		if err != nil {
			return errors.Wrap(err, "failed to send message")
		}
	}

	return nil
}
