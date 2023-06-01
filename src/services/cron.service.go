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
			// Create Object interface
			scraper := interfaces.Scraper{
				AllowedDomains: []string{constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN},
			}

			// Create Collector
			collector := scraper.InitCollector()
			scrapedProduct, err := scrapers.SendAmazonCollyRequest(URL.String(), scraper, collector)
			if err != nil {
				log.Println("Error Getting the data from the craping ", err)
			}

			// Create the product and store in DB
			err = database.CreateProduct(scrapedProduct, "")
			if err != nil {
				log.Println("Error creating the Amazon product")
			}

			continue
		}

		if strings.Contains(URL.Host, "exito") {
			// Make the Scraping to the page
			scrapedProduct, err := scrapers.SendExitoCollyRequest(URL.String())
			if err != nil {
				log.Println("Error Getting the data from the craping ", err)
			}

			// Create the product and store in DB
			err = database.CreateProduct(scrapedProduct, "")
			if err != nil {
				log.Println("Error creating the Amazon product")
			}

			continue
		}

		if strings.Contains(URL.Host, "jumbo") {
			// Make the Scraping to the page
			scrapedProduct, err := scrapers.SendJumboCollyRequest(URL.String())
			if err != nil {
				log.Println("Error Getting the data from the craping ", err)
			}

			// Create the product and store in DB
			err = database.CreateProduct(scrapedProduct, "")
			if err != nil {
				log.Println("Error creating the Amazon product")
			}

			continue
		}
	}

}
