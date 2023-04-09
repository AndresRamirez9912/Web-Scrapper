package scrapers

import (
	"log"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/models"

	"github.com/gocolly/colly"
)

var jumboData = &models.JumboProduct{
	Description: make(map[int]string),
}

func SendJumboCollyRequest(productURL string) (*models.JumboProduct, error) {
	// Clean the objData
	jumboData.Name = ""

	// Create a collector to setup the data searcher
	scraper := interfaces.Scraper{
		AllowedDomains: []string{constants.JUMBO_HALF_DOMAIN, constants.JUMBO_DOMAIN},
	}

	collector := scraper.InitCollector()

	// Callbacks
	collector.OnError(scraper.OnError)

	collector.OnRequest(scraper.OnRequest)

	collector.OnResponse(scraper.OnResponse)

	collector.OnHTML(constants.JUMBO_QUERY_SELECTOR, jumboOnHTML)

	// Visit the page
	err := collector.Visit(productURL)
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}

	collector.Wait()

	scrapedElement, err := jumboHandleResponse()
	if err != nil {
		log.Fatal("Error getting data from scraping")
		return nil, err
	}

	return scrapedElement, nil
}

func jumboOnHTML(h *colly.HTMLElement) {
	if jumboData.Name == "" {
		jumboData.Name = h.ChildText(constants.JUMBO_QUERY_NAME)
		jumboData.ImageURL = h.ChildAttr(constants.JUMBO_QUERY_IMAGE_URL, "src")
		jumboData.Description[0] = h.ChildText(constants.JUMBO_QUERY_DESCRIPTION)
		jumboData.CurrentPrice = h.ChildText(constants.JUMBO_QUERY_CURRENTPRICE)
		jumboData.HighPrice = h.ChildText(constants.JUMBO_QUERY_HIGHTPRICE)
		jumboData.Disccount = h.ChildText(constants.JUMBO_QUERY_DISCOUNT)
	}
}

func jumboHandleResponse() (*models.JumboProduct, error) {
	return jumboData, nil
}
