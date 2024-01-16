package customermailing

import "time"

type Message struct {
	Email      string
	Title      string
	Content    string
	InsertTime time.Time
	MailingID  int
}
