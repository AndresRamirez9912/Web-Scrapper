package services

import (
	"log"
	"net/url"
	"strings"
	"webScraper/src/constants"
	"webScraper/src/database"
	"webScraper/src/interfaces"
	"webScraper/src/services/scrapers"
)

func CheckProducts() {
	//TODO Optimize with Goroutines
	// Get the URL of the products
	products, err := database.GetAllProducts()
	if err != nil {
		log.Println("Error Getting the products")
		return
	}

	for _, productURL := range products {
		// Get the host
		URL, err := url.Parse(productURL)
		if err != nil {
			log.Println("Error Parsing the URL")
		}

		// Check the product based on the host
		if strings.Contains(URL.Host, "amazon") {
			// Make the Scraping to the page

			// Create Collector
			amazonScraper := interfaces.ColectorFactory("amazon")
			err = scrapers.ScrapedPage(URL.String(), []string{constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN}, amazonScraper)
			if err != nil {
				log.Println("Error Getting the data from the craping ", err)
			}

			// Create the product and store in DB
			err = database.CreateProduct(amazonScraper, "")
			if err != nil {
				log.Println("Error creating the Amazon product")
			}

			continue
		}

		if strings.Contains(URL.Host, "exito") {
			// Make the Scraping to the page
			exitoScraper := interfaces.ColectorFactory("exito")
			err = scrapers.ScrapedPage(URL.String(), []string{constants.EXITO_HALF_DOMAIN, constants.EXITO_DOMAIN}, exitoScraper)
			if err != nil {
				log.Println("Error Getting the data from the craping ", err)
			}

			// Create the product and store in DB
			err = database.CreateProduct(exitoScraper, "")
			if err != nil {
				log.Println("Error creating the Amazon product")
			}

			continue
		}

		if strings.Contains(URL.Host, "jumbo") {
			// Make the Scraping to the page
			jumboScraper := interfaces.ColectorFactory("jumbo")
			err = scrapers.ScrapedPage(URL.String(), []string{constants.JUMBO_HALF_DOMAIN, constants.JUMBO_DOMAIN}, jumboScraper)
			if err != nil {
				log.Println("Error Getting the data from the craping ", err)
			}

			// Create the product and store in DB
			err = database.CreateProduct(jumboScraper, "")
			if err != nil {
				log.Println("Error creating the Amazon product")
			}

			continue
		}
	}

}
