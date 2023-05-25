package interfaces

import "gopkg.in/gomail.v2"

type Senders interface {
	DialAndSend(m ...*gomail.Message) error
}
