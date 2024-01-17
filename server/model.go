package server

import "time"

type MailingMessage struct {
	Email      string `json:"email"`
	Title      string
	Content    string
	InsertTime time.Time
	MailingID  int
}
