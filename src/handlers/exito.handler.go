package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"webScraper/src/models"
	"webScraper/src/pages"
)

func GetExitoData(w http.ResponseWriter, r *http.Request) {
	// Get the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error receiving the request ", err)
	}

	// Convert the data into an object
	product := &models.ExitoResponse{}
	err = json.Unmarshal(body, product)
	if err != nil {
		log.Fatal("Error Getting the data from the body request ", err)
	}

	// Make the Scraping to the page
	scrapedProduct, err := sendCollyRequest(product.ProductURL)
	if err != nil {
		log.Fatal("Error Getting the data from the craping ", err)
	}

	// Send the response
	dataResponse, err := json.Marshal(scrapedProduct)
	if err != nil {
		log.Fatal("Error Serializing the obtained data ", err)
	}

	w.Header().Set("Content-Type", "application/json")
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

	collector.OnHTML("script[type='application/ld+json']", pages.ExitoOnHTML)

	// Visit the page
	err := collector.Visit(productURL)
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}

	collector.Wait()

	exitoData, err := pages.HandleResponse()
	if err != nil {
		log.Fatal("Error getting data from scraping")
		return nil, err
	}

	return exitoData, nil
}
