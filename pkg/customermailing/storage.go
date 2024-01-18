package customermailing

import (
	"context"
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

func (s *Storage) DeleteOlderThan(ctx context.Context, mailingID int, olderThan time.Time) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM mailing_messages WHERE mailing_id = $1 AND insert_time <= $2", mailingID, olderThan)
	return err
}

func (s *Storage) GetMailingMessagesByID(ctx context.Context, id int) ([]MailingMessage, error) {
	res := []MailingMessage{}
	err := s.db.SelectContext(ctx, &res, "SELECT email, title, content, insert_time, mailing_id FROM mailing_messages WHERE mailing_id = $1", id)
	return res, err
}
