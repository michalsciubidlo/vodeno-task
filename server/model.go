package server

import "time"

type MailingMessage struct {
	Email      string    `json:"email"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	InsertTime time.Time `json:"insert_time"`
	MailingID  int       `json:"mailing_id"`
}
