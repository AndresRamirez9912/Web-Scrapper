package scrapers

import (
	"errors"
	"log"
	"strings"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/models/scraping"

	"github.com/gocolly/colly"
)

var amazonData = &scraping.AmazonProduct{}

func SendAmazonCollyRequest(productURL string) (*scraping.AmazonProduct, error) {

	// Create Object interface
	scraper := interfaces.Scraper{
		AllowedDomains: []string{constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN},
	}

	// Create Collector
	collector := scraper.InitCollector()

	// Callbacks
	collector.OnError(scraper.OnError)

	collector.OnRequest(scraper.OnRequest)

	collector.OnResponse(scraper.OnResponse)

	collector.OnHTML(constants.AMAZON_QUERY_SELECTOR, amazonOnHTML)

	// Visit the page
	err := collector.Visit(productURL)
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}

	collector.Wait()

	data, err := handleResponse()
	if err != nil {
		log.Fatal("Error getting data from scraping")
		return nil, err
	}

	// Get the Product Id from the URL
	productId, err := getProductId(productURL)
	if err != nil {
		log.Fatal("Error getting the Id of the product")
		return nil, err
	}

	data.Id = productId
	data.ProductURL = productURL
	return data, nil
}

func amazonOnHTML(h *colly.HTMLElement) {
	amazonData.Name = h.ChildText(constants.AMAZON_QUERY_NAME)   // Product Tittle
	amazonData.Brand = h.ChildText(constants.AMAZON_QUERY_BRAND) // Brand

	// Description
	var description = make(map[int]string)
	h.ForEach(constants.AMAZON_QUERY_DESCRIPTION, func(i int, h *colly.HTMLElement) {
		description[i] = h.Text
	})
	amazonData.Description = description

	if h.ChildAttr(constants.AMAZON_QUERY_IMAGE_URL, constants.AMAZON_QUERY_IMAGE_URL_ATTR) != "" {
		amazonData.ImageURL = h.ChildAttr(constants.AMAZON_QUERY_IMAGE_URL, constants.AMAZON_QUERY_IMAGE_URL_ATTR) // Image URL
	}

	// Discount Form
	amazonData.Disccount = h.ChildText(constants.AMAZON_QUERY_DISCOUNT_DISCOUNT)        // Product Discount
	amazonData.CurrentPrice = h.ChildText(constants.AMAZON_QUERY_CURRENTPRICE_DISCOUNT) // Product Lower Price
	amazonData.HighPrice = h.ChildText(constants.AMAZON_QUERY_HIGHTPRICE_DISCOUNT)      // Original Price, withou Discount

	// Current Form - No discount
	if amazonData.HighPrice == "" {
		amazonData.HighPrice = h.ChildText(constants.AMAZON_QUERY_HIGHTPRICE_CURRENT)
		amazonData.CurrentPrice = h.ChildText(constants.AMAZON_QUERY_CURRENTPRICE_CURRENT)
		prices := strings.Split(h.ChildText(constants.AMAZON_QUERY_HIGHTPRICE_CURRENT), "US")
		if len(prices) > 2 {
			amazonData.HighPrice = prices[2]
			amazonData.CurrentPrice = prices[2]
		}
	}

	// Table Form
	if amazonData.HighPrice == "" {
		var prices = make(map[int]string)
		h.ForEach(constants.AMAZON_QUERY_PRICES_TABLE, func(i int, h *colly.HTMLElement) {
			prices[i] = h.ChildText(constants.AMAZON_QUERY_PRICE_ELEMENT_TABLE)
		})
		amazonData.HighPrice = prices[0]
		amazonData.CurrentPrice = prices[1]
		amazonData.Disccount = prices[2]
	}

}

func handleResponse() (*scraping.AmazonProduct, error) {
	return amazonData, nil
}

func getProductId(productURL string) (string, error) {
	urlElements := strings.Split(productURL, "/")

	for index, value := range urlElements {
		if value == "dp" {
			return urlElements[index+1], nil
		}
	}
	return "", errors.New("Product Id not found")
}
