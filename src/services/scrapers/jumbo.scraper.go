package scrapers

import (
	"log"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/models/scraping"

	"github.com/gocolly/colly"
)

var jumboData = &scraping.JumboProduct{}

func SendJumboCollyRequest(productURL string) (*scraping.JumboProduct, error) {
	// Clean the objData
	jumboData = &scraping.JumboProduct{}

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
		log.Println("Error Visiting the page ", err)
	}

	collector.Wait()

	scrapedElement, err := jumboHandleResponse()
	if err != nil {
		log.Println("Error getting data from scraping")
		return nil, err
	}

	// Get the Last Information
	scrapedElement.ProductURL = productURL

	return scrapedElement, nil
}

func jumboOnHTML(h *colly.HTMLElement) {
	if jumboData.Name == "" {
		jumboData.Id = h.ChildText(constants.JUMBO_QUERY_PRODUCT_ID)
		jumboData.Name = h.ChildText(constants.JUMBO_QUERY_NAME)
		jumboData.ImageURL = h.ChildAttr(constants.JUMBO_QUERY_IMAGE_URL, "src")
		jumboData.Description = h.ChildText(constants.JUMBO_QUERY_DESCRIPTION)
		jumboData.CurrentPrice = h.ChildText(constants.JUMBO_QUERY_CURRENTPRICE)
		jumboData.HighPrice = h.ChildText(constants.JUMBO_QUERY_HIGHTPRICE)
		jumboData.Disccount = h.ChildText(constants.JUMBO_QUERY_DISCOUNT)
	}
}

func jumboHandleResponse() (*scraping.JumboProduct, error) {
	return jumboData, nil
}
