package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/services/scrapers"
	"webScraper/src/utils"
)

func GetJumboData(w http.ResponseWriter, r *http.Request) {
	// Get the User Id from the Cookie
	userId := r.Context().Value(constants.USER_ID).(string)

	// Get the URL of the product
	URL, err := utils.GetProductURL(r)
	if err != nil {
		log.Fatal("Error Getting the URL of the Product ", err)
	}

	// Make the Scraping to the page
	scrapedProduct, err := scrapers.SendJumboCollyRequest(URL)
	if err != nil {
		log.Fatal("Error Getting the data from the craping ", err)
	}

	// Send the response
	dataResponse, err := json.Marshal(scrapedProduct)
	if err != nil {
		log.Fatal("Error Serializing the obtained data ", err)
	}

	// Create the product in the DB
	err = interfaces.CreateProduct(scrapedProduct, userId)
	if err != nil {
		log.Println("Error creating the Amazon product")
	}

	w.Header().Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	w.WriteHeader(http.StatusOK) // Wite status and previous headers into request
	w.Write(dataResponse)        // Send body
}
