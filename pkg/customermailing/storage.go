package customermailing

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Add(ctx context.Context, msg MailingMessage) error {
	_, err := s.db.NamedExecContext(ctx, "INSERT INTO mailing_messages(email, title, content, insert_time, mailing_id) VALUES (:email, :title, :content, :insert_time, :mailing_id)", msg)
	return err
}

func (s *Storage) DeleteOlderThan(ctx context.Context, t time.Time) error {
	return fmt.Errorf("storage.DeleteOlderThan: not implemented")
}

func (s *Storage) GetMailingMessagesByID(ctx context.Context, id int) ([]MailingMessage, error) {
	return nil, fmt.Errorf("storage.GetMailingMessagesByID: not implemented")
}
