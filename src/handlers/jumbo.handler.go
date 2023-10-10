package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webScraper/src/constants"
	"webScraper/src/database"
	"webScraper/src/models/scraping"
	"webScraper/src/services/scrapers"
	"webScraper/src/utils"
)

func GetJumboData(w http.ResponseWriter, r *http.Request) {
	// Get the User Id from the Cookie
	userId := r.Context().Value(constants.USER_ID).(string)

	// Get the URL of the product
	URL, err := utils.GetProductURL(r)
	if err != nil {
		log.Println("Error Getting the URL of the Product ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Make the Scraping to the page
	jumboScraper := &scraping.JumboScraper{}
	err = scrapers.ScrapedPage(URL, []string{constants.JUMBO_HALF_DOMAIN, constants.JUMBO_DOMAIN}, jumboScraper)
	if err != nil {
		log.Println("Error Getting the data from the craping ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	dataResponse, err := json.Marshal(jumboScraper)
	if err != nil {
		log.Println("Error Serializing the obtained data ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create the product in the DB
	err = database.CreateProduct(jumboScraper, userId)
	if err != nil {
		log.Println("Error creating the Amazon product")
		w.WriteHeader(http.StatusInternalServerError)
		database.DeleteProduct(jumboScraper)
		return
	}

	w.Header().Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)   // Wite status and previous headers into request
	_, err = w.Write(dataResponse) // Send body
	if err != nil {
		log.Println("Error sending the response ", err)
		w.WriteHeader(http.StatusInternalServerError)
		database.DeleteProduct(jumboScraper)
		return
	}
}
