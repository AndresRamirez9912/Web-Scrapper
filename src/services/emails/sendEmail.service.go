package services

import (
	"bytes"
	"log"
	"net/url"
	"os"
	"strconv"
	"text/template"
	"time"
	"webScraper/src/constants"
	"webScraper/src/models/auth"

	"gopkg.in/gomail.v2"
)

func sendEmail(toEmail string, subject string, body string) error {
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

func SendVerificationEmail(user *auth.User) error {
	var body bytes.Buffer

	// Get the Template with the values
	template, err := template.ParseFiles("src/services/emails/templates/emailVerification.template.html")
	if err != nil {
		log.Fatal("Error Trying to get the template ", err)
		return err
	}

	// Create the struct with the data to send to the template
	url, err := url.Parse("http://localhost:3000/verify")
	if err != nil {

	}
	query := url.Query()
	query.Set("user", user.Id)
	query.Set("time", strconv.FormatInt(time.Now().Unix(), 10))
	url.RawQuery = query.Encode()

	data := struct {
		UserName string
		Link     string
	}{
		UserName: user.Name,
		Link:     url.String(),
	}

	// Execute the template and get the string
	err = template.Execute(&body, data)
	if err != nil {
		log.Fatal("Error Trying to execute the template ", err)
		return err
	}

	// Send the email
	err = sendEmail(user.Email, constants.ACCOUNT_VERIFICATION_SUBJECT, body.String())
	if err != nil {
		log.Fatal("Error Sending the email", err)
		return err
	}
	return nil
}
