package handlers

import (
	"log"
	"net/http"
	"webScraper/src/interfaces"
	"webScraper/src/services/scrapers"
	"webScraper/src/utils"
)

func GetExitoData(w http.ResponseWriter, r *http.Request) {
	// Get the URL of the product
	URL, err := utils.GetProductURL(r)
	if err != nil {
		log.Println("Error Getting the URL of the Product ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Create Collector
	exitoScraper := interfaces.ColectorFactory("exito")
	err = scrapers.ScrapedPage(URL, exitoScraper)
	if err != nil {
		log.Println("Error Getting the data from the craping ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Store in the DB and send the response
	err = utils.StoreAndSendResponse(w, r, exitoScraper)
	if err != nil {
		return
	}

}
