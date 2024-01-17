package customermailing

import "time"

type MailingMessage struct {
	Email      string
	Title      string
	Content    string
	InsertTime time.Time
	MailingID  int
}
