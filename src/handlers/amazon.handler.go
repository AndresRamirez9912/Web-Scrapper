package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webScraper/src/constants"
	"webScraper/src/models"
	"webScraper/src/pages"
	"webScraper/src/utils"
)

func GetAmazonData(w http.ResponseWriter, r *http.Request) {
	// Get the URL of the product
	URL, err := utils.GetProductURL(r)
	if err != nil {
		log.Fatal("Error Getting the URL of the Product ", err)
	}

	// Make the Scraping to the page
	scrapedProduct, err := sendColly(URL)
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

func sendColly(productURL string) (*models.AmazonProduct, error) {

	// Create a collector to setup the data searcher
	collector := pages.InitAmazonCollector()

	// Callbacks
	collector.OnError(pages.AmazonOnError)

	collector.OnRequest(pages.AmazonOnRequest)

	collector.OnResponse(pages.AmazonOnResponse)

	collector.OnHTML("div#centerCol.centerColAlign, li[data-csa-c-action='image-block-main-image-hover']", pages.AmazonOnHTML)

	// Visit the page
	err := collector.Visit(productURL)
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}

	collector.Wait()

	data, err := pages.AmazonHandleResponse()
	if err != nil {
		log.Fatal("Error getting data from scraping")
		return nil, err
	}
	fmt.Println("Finish Scraping")
	return data, nil
}
