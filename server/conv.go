package server

import "github.com/michalsciubidlo/vodeno-task/pkg/customermailing"

func toCustomerMailingMessage(m MailingMessage) customermailing.MailingMessage {
	return customermailing.MailingMessage{
		Email:      m.Email,
		Title:      m.Title,
		Content:    m.Content,
		InsertTime: m.InsertTime,
		MailingID:  m.MailingID,
	}
}
