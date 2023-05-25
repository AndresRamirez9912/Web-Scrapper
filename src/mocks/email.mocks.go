package mocks

import (
	"errors"

	"gopkg.in/gomail.v2"
)

type MockGomail struct {
	Success bool
}

func (mock *MockGomail) DialAndSend(m ...*gomail.Message) error {
	if mock.Success {
		return nil
	}
	return errors.New("The Email could not be sent")
}
