package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webScraper/src/constants"
	"webScraper/src/models"
	"webScraper/src/pages"
	"webScraper/src/utils"
)

func GetExitoData(w http.ResponseWriter, r *http.Request) {
	// Get the URL of the product
	URL, err := utils.GetProductURL(r)
	if err != nil {
		log.Fatal("Error Getting the URL of the Product ", err)
	}

	// Make the Scraping to the page
	scrapedProduct, err := sendCollyRequest(URL)
	if err != nil {
		log.Fatal("Error Getting the data from the craping ", err)
	}

	// Send the response
	dataResponse, err := json.Marshal(scrapedProduct)
	if err != nil {
		log.Fatal("Error Serializing the obtained data ", err)
	}

	w.Header().Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	w.WriteHeader(http.StatusOK) // Wite status and previous headers into request
	w.Write(dataResponse)        // Send body

}

func sendCollyRequest(productURL string) (*models.ExitoProduct, error) {

	// Create a collector to setup the data searcher
	collector := pages.InitExitoCollector()

	// Callbacks
	collector.OnError(pages.ExitoOnError)

	collector.OnRequest(pages.ExitoOnRequest)

	collector.OnResponse(pages.ExitoOnResponse)

	collector.OnHTML(constants.EXITO_QUERY_SELECTOR, pages.ExitoOnHTML)

	// Visit the page
	err := collector.Visit(productURL)
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}

	collector.Wait()

	exitoData, err := pages.ExitoHandleResponse()
	if err != nil {
		log.Fatal("Error getting data from scraping")
		return nil, err
	}

	return exitoData, nil
}
