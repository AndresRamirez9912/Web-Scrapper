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
	"webScraper/src/interfaces"
	"webScraper/src/models/auth"
	"webScraper/src/models/scraping"

	"gopkg.in/gomail.v2"
)

func sendEmail(toEmail string, subject string, body string, sender interfaces.Senders) error {
	// Create the message
	msg := gomail.NewMessage()
	msg.SetHeader(constants.FROM_HEADER, os.Getenv(constants.MY_EMAIL))
	msg.SetHeader(constants.TO_HEADER, toEmail)
	msg.SetHeader(constants.SUBJECT_HEADER, subject)
	msg.SetBody(constants.CONTENT_TYPE_EMAIL, body)

	// Send the Email
	err := sender.DialAndSend(msg)
	if err != nil {
		log.Println("Error sending the email ", err)
		return err
	}
	return nil
}

func SendVerificationEmail(user *auth.User, sender interfaces.Senders) error {
	var body bytes.Buffer

	// Get the Template with the values
	template, err := getTemplate(constants.VERIFICATION_TEMPLATE_ABSOLUTE_ADDRESS, constants.VERIFICATION_TEMPLATE_RELATIVE_ADDRESS)
	if err != nil {
		log.Println("Error Trying to get the template ", err)
		return err
	}

	// Create the struct with the data to send to the template
	url, err := url.Parse(constants.VERIFICATION_URL)
	if err != nil {
		log.Println("Error Creating the Verification Link ", err)
		return err
	}
	query := url.Query()
	query.Set(constants.USER_QUERY, user.Id)
	query.Set(constants.TIME_QUERY, strconv.FormatInt(time.Now().Unix(), 10))
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
		log.Println("Error Trying to execute the template ", err)
		return err
	}

	// Send the email
	err = sendEmail(user.Email, constants.ACCOUNT_VERIFICATION_SUBJECT, body.String(), sender)
	if err != nil {
		log.Println("Error Sending the email", err)
		return err
	}
	return nil
}

func SendNotificationLowerPrice(user *auth.User, sender interfaces.Senders, product scraping.Product) error {
	var body bytes.Buffer

	// Get the Template
	template, err := getTemplate(constants.LOWER_PRICE_TEMPLATE_ABSOLUTE_ADDRESS, constants.LOWER_PRICE_TEMPLATE_RELATIVE_ADDRESS)
	if err != nil {
		log.Println("Error Trying to get the template ", err)
		return err
	}

	// Assign the values
	err = template.Execute(&body, product)
	if err != nil {
		log.Println("Error Trying to execute the template ", err)
		return err
	}

	// Send the email
	err = sendEmail(user.Email, constants.LOWER_NOTIFICATION_SUBJECT, body.String(), sender)
	if err != nil {
		log.Println("Error Sending the email", err)
		return err
	}
	log.Println("Lower Price alert sent")
	return nil
}

func getTemplate(relativeURL string, absoluteURL string) (*template.Template, error) {
	_, err := os.Stat("go.mod") // Check if this file exists
	// I'm in the root position
	if err != nil {
		return template.ParseFiles(absoluteURL)

	}
	// I'm in the services position
	return template.ParseFiles(relativeURL)
}
