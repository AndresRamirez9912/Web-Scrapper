package interfaces

import (
	"webScraper/src/models/scraping"

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
	CreateProductStructure(userId string) scraping.Product
}

func ColectorFactory(collectorType string) Collectors {
	switch collectorType {
	case "amazon":
		return &scraping.AmazonProduct{}
	case "jumbo":
		return &scraping.JumboScraper{}
	case "exito":
		return &scraping.ExitoScraper{}
	}
	return nil
}
