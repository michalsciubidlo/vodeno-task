package customermailing

import "time"

type MailingMessage struct {
	Email      string    `db:"email"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	InsertTime time.Time `db:"insert_time"`
	MailingID  int       `db:"mailing_id"`
}
