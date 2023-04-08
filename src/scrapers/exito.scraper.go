package scrapers

import (
	"encoding/json"
	"log"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/models"

	"github.com/gocolly/colly"
)

var exitoData string

func SendExitoCollyRequest(productURL string) (*models.ExitoProduct, error) {

	// Create a collector to setup the data searcher
	exitoScraper := interfaces.Scraper{
		AllowedDomains: []string{constants.EXITO_HALF_DOMAIN, constants.EXITO_DOMAIN},
	}

	collector := exitoScraper.InitCollector()

	// Callbacks
	collector.OnError(exitoScraper.OnError)

	collector.OnRequest(exitoScraper.OnRequest)

	collector.OnResponse(exitoScraper.OnResponse)

	collector.OnHTML(constants.EXITO_QUERY_SELECTOR, onHTML)

	// Visit the page
	err := collector.Visit(productURL)
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}

	collector.Wait()

	exitoData, err := exitoHandleResponse()
	if err != nil {
		log.Fatal("Error getting data from scraping")
		return nil, err
	}

	return exitoData, nil
}

func onHTML(h *colly.HTMLElement) {
	exitoData = h.Text // Send the response
}

func exitoHandleResponse() (*models.ExitoProduct, error) {
	exitoProduct := &models.ExitoProduct{}
	err := json.Unmarshal([]byte(exitoData), exitoProduct)
	if err != nil {
		log.Fatal("Error unmarshaling scraping response ", err)
		return nil, err
	}
	return exitoProduct, nil
}
