package services

import (
	"log"
	"os"
	"webScraper/src/constants"

	"gopkg.in/gomail.v2"
)

func SendEmail(toEmail string, subject string, body string) error {
	// Create the message
	msg := gomail.NewMessage()
	msg.SetHeader(constants.FROM_HEADER, constants.MY_EMAIL)
	msg.SetHeader(constants.TO_HEADER, toEmail)
	msg.SetHeader(constants.SUBJECT_HEADER, subject)
	msg.SetBody(constants.CONTENT_TYPE_EMAIL, body)

	// Send the Email
	sender := gomail.NewDialer(constants.SMTP_HOST, 587, constants.MY_EMAIL, os.Getenv(constants.EMAIL_PASSWORD))
	err := sender.DialAndSend(msg)
	if err != nil {
		log.Panic("Error sending the email ", err)
		return err
	}
	return nil
}
