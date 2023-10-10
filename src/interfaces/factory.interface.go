package interfaces

import (
	"webScraper/src/constants"
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
	GetDomains() []string
}

func ColectorFactory(collectorType string) Collectors {
	switch collectorType {
	case "amazon":
		return &scraping.AmazonColector{AllowedDomains: []string{constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN}}
	case "jumbo":
		return &scraping.JumboCollector{AllowedDomains: []string{constants.JUMBO_HALF_DOMAIN, constants.JUMBO_DOMAIN}}
	case "exito":
		return &scraping.ExitoCollector{AllowedDomains: []string{constants.EXITO_HALF_DOMAIN, constants.EXITO_DOMAIN}}
	}
	return nil
}
