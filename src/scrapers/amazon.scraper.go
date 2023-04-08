package scrapers

import (
	"fmt"
	"log"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/models"

	"github.com/gocolly/colly"
)

var amazonData = &models.AmazonProduct{}

func SendAmazonCollyRequest(productURL string) (*models.AmazonProduct, error) {

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

	collector.OnHTML("div#centerCol.centerColAlign, li[data-csa-c-action='image-block-main-image-hover']", amazonOnHTML)

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
	fmt.Println("Finish Scraping")
	return data, nil
}

func amazonOnHTML(h *colly.HTMLElement) {
	amazonData.Name = h.ChildText("span[id='productTitle']")                                      // Product Tittle
	amazonData.Brand = h.ChildText("div.celwidget  tbody tr.a-spacing-small.po-brand td.a-span9") // Brand

	// Description
	var description = make(map[int]string)
	h.ForEach("div[id='feature-bullets'] ul li span.a-list-item", func(i int, h *colly.HTMLElement) {
		description[i] = h.Text
	})
	amazonData.Description = description

	if h.ChildAttr("img.a-dynamic-image", "data-old-hires") != "" {
		amazonData.ImageURL = h.ChildAttr("img.a-dynamic-image", "data-old-hires") // Image URL
	}

	// Discount Form
	amazonData.Disccount = h.ChildText("div.a-section.a-spacing-none.aok-align-center span.a-size-large.a-color-price.savingPriceOverride") // Product Discount
	amazonData.CurrentPrice = h.ChildText("div.a-section.a-spacing-none.aok-align-center span.a-offscreen")                                 // Product Lower Price
	amazonData.HighPrice = h.ChildText("div.a-section.a-spacing-small.aok-align-center span.a-offscreen")                                   // Original Price, withou Discount

	// Current Form
	if amazonData.HighPrice == "" {
		amazonData.HighPrice = h.ChildText(".a-section.a-spacing-none.aok-align-center span.a-offscreen")
		amazonData.CurrentPrice = h.ChildText(".a-section.a-spacing-none.aok-align-center span.a-offscreen")
	}

	// Table Form
	if amazonData.HighPrice == "" {
		var prices = make(map[int]string)
		h.ForEach("div.a-section.a-spacing-small table.a-lineitem.a-align-top tr", func(i int, h *colly.HTMLElement) {
			prices[i] = h.ChildText("span.a-price.a-text-price.a-size-base span.a-offscreen ,span.a-price.a-text-price.a-size-medium.apexPriceToPay span.a-offscreen")
		})
		amazonData.HighPrice = prices[0]
		amazonData.CurrentPrice = prices[1]
		amazonData.Disccount = prices[2]
	}

}

func handleResponse() (*models.AmazonProduct, error) {
	return amazonData, nil
}
