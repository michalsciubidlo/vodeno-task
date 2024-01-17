package customermailing

import (
	"context"
	"fmt"
	"time"
)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Add(ctx context.Context, msg MailingMessage) error {
	return fmt.Errorf("storage.Add: not implemented")
}

func (s *Storage) DeleteOlderThan(ctx context.Context, t time.Time) error {
	return fmt.Errorf("storage.DeleteOlderThan: not implemented")
}

func (s *Storage) GetMailingMessagesByID(ctx context.Context, id int) ([]MailingMessage, error) {
	return nil, fmt.Errorf("storage.GetMailingMessagesByID: not implemented")
}
