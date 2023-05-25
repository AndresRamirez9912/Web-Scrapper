package services

import (
	"testing"
	"webScraper/src/mocks"
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
