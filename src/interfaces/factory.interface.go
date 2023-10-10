package interfaces

import (
	"github.com/gocolly/colly"
)

// This is the super interface that every object will implement
type Collectors interface {
	OnError(*colly.Response, error)
	OnRequest(*colly.Request)
	OnResponse(*colly.Response)
	OnHTML(*colly.HTMLElement)
	GetQuerySelector() string
	SetURL(string)
}
