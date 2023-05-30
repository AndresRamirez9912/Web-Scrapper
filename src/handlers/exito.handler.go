package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webScraper/src/constants"
	"webScraper/src/database"
	"webScraper/src/services/scrapers"
	"webScraper/src/utils"
)

func GetExitoData(w http.ResponseWriter, r *http.Request) {
	// Get the User Id from the Cookie
	userId := r.Context().Value(constants.USER_ID).(string)

	// Get the URL of the product
	URL, err := utils.GetProductURL(r)
	if err != nil {
		log.Println("Error Getting the URL of the Product ", err)
	}

	// Make the Scraping to the page
	scrapedProduct, err := scrapers.SendExitoCollyRequest(URL)
	if err != nil {
		log.Println("Error Getting the data from the craping ", err)
	}

	// Send the response
	dataResponse, err := json.Marshal(scrapedProduct)
	if err != nil {
		log.Println("Error Serializing the obtained data ", err)
	}

	// Create the product and store in DB
	err = database.CreateProduct(scrapedProduct, userId)
	if err != nil {
		log.Println("Error creating the Amazon product")
	}

	w.Header().Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)   // Wite status and previous headers into request
	_, err = w.Write(dataResponse) // Send body
	if err != nil {
		log.Println("Error sending the response ", err)
		return
	}

}
