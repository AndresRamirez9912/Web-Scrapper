package scrapers

import (
	"log"
	"webScraper/src/interfaces"
	"webScraper/src/utils"
)

func ScrapedPage(productURL string, scraper interfaces.Collectors) error {
	// Create the Colly scraper
	collector := utils.InitCollector(scraper.GetDomains())

	// Save URL in my object
	scraper.SetURL(productURL)

	// Callbacks
	collector.OnError(scraper.OnError)

	collector.OnRequest(scraper.OnRequest)

	collector.OnResponse(scraper.OnResponse)

	collector.OnHTML(scraper.GetQuerySelector(), scraper.OnHTML)

	// Visit the page
	err := collector.Visit(productURL)
	if err != nil {
		log.Println("Error Visiting the page ", err)
		return err
	}

	// Wait meanwhile the page is cratched
	collector.Wait()
	return nil
}
