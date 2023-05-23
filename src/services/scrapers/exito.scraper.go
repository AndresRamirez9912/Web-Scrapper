package scrapers

import (
	"encoding/json"
	"log"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/models/scraping"

	"github.com/gocolly/colly"
)

var exitoProduct = &scraping.ExitoProduct{}
var exitoData string

func SendExitoCollyRequest(productURL string) (*scraping.ExitoProduct, error) {
	// Clear the Object
	exitoProduct = &scraping.ExitoProduct{}
	exitoData = ""

	// Create a collector to setup the data searcher
	exitoScraper := interfaces.Scraper{
		AllowedDomains: []string{constants.EXITO_HALF_DOMAIN, constants.EXITO_DOMAIN},
	}

	collector := exitoScraper.InitCollector()

	// Callbacks
	collector.OnError(exitoScraper.OnError)

	collector.OnRequest(exitoScraper.OnRequest)

	collector.OnResponse(exitoScraper.OnResponse)

	collector.OnHTML(constants.EXITO_QUERY_SELECTOR, exitoOnHTML)

	// Visit the page
	err := collector.Visit(productURL)
	if err != nil {
		log.Println("Error Visiting the page ", err)
		return nil, err
	}

	collector.Wait()

	exitoData, err := exitoHandleResponse()
	if err != nil {
		log.Println("Error getting data from scraping")
		return nil, err
	}

	return exitoData, nil
}

func exitoOnHTML(h *colly.HTMLElement) {
	if exitoData == "" {
		exitoData = h.ChildText("script[type='application/ld+json']") // Send the response
	}
	exitoProduct.Id = h.ChildText("span.vtex-product-identifier-0-x-product-identifier__value")
}

func exitoHandleResponse() (*scraping.ExitoProduct, error) {
	err := json.Unmarshal([]byte(exitoData), exitoProduct)
	if err != nil {
		log.Println("Error unmarshaling scraping response ", err)
		return nil, err
	}
	return exitoProduct, nil
}
