package services

import (
	"testing"
	"webScraper/src/mocks"
	"webScraper/src/models/auth"
)

func TestSendEmail(t *testing.T) {
	t.Run("Success Test", func(t *testing.T) {
		gomailMock := &mocks.MockGomail{Success: true}
		err := sendEmail("a@gmail.com", "Testing", "Hello I'm a test", gomailMock)
		if err != nil {
			t.Error("Error")
		}
	})

	t.Run("Error sending the Email", func(t *testing.T) {
		gomailMock := &mocks.MockGomail{Success: false}
		err := sendEmail("a@gmail.com", "Testing", "Hello I'm a test", gomailMock)
		if err == nil {
			t.Error("Error")
		}
	})
}

func TestSendVerificationEmail(t *testing.T) {
	t.Run("Success Test", func(t *testing.T) {
		// Create the user object
		user := auth.User{
			Name:     "",
			Password: "",
			Email:    "",
		}

		// Execute the function
		gomailMock := &mocks.MockGomail{Success: true}
		err := SendVerificationEmail(&user, gomailMock)
		if err != nil {
			t.Error("Error sending the email")
		}
	})

	t.Run("Fail Case, Error sending the email", func(t *testing.T) {
		// Create the user object
		user := auth.User{
			Name:     "",
			Password: "",
			Email:    "",
		}

		// Execute the function
		gomailMock := &mocks.MockGomail{Success: false}
		err := SendVerificationEmail(&user, gomailMock)
		if err == nil {
			t.Error("Error sending the email")
		}
	})
}
