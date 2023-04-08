package scrapers

import (
	"fmt"
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

	collector.OnHTML("div.vtex-store-components-3-x-productImage, div.vtex-flex-layout-0-x-flexCol.ml0.mr0.pl0.pr0.flex.flex-column.h-100.w-100", jumboOnHTML)

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
	// Name
	if jumboData.Name == "" {
		jumboData.Name = h.ChildText("span.vtex-store-components-3-x-productBrand.vtex-store-components-3-x-productBrand--quickview")
		jumboData.ImageURL = h.ChildAttr("img.vtex-store-components-3-x-productImageTag.vtex-store-components-3-x-productImageTag--main", "src")
		jumboData.Description[0] = h.ChildText("div.vtex-store-components-3-x-content.h-auto")
		jumboData.CurrentPrice = h.ChildText("div#items-price.flex.c-emphasis.tiendasjumboqaio-jumbo-minicart-2-x-cencoPrice div.tiendasjumboqaio-jumbo-minicart-2-x-price")
		jumboData.HighPrice = h.ChildText("div.b.ml2.tiendasjumboqaio-jumbo-minicart-2-x-cencoListPriceWrapper div.tiendasjumboqaio-jumbo-minicart-2-x-price")
		jumboData.Disccount = h.ChildText("div.pr7.items-stretch.vtex-flex-layout-0-x-stretchChildrenWidth.flex span.vtex-product-price-1-x-currencyContainer.vtex-product-price-1-x-currencyContainer--summary")
	}
	fmt.Printf("%+v\n\n", jumboData)
}

func jumboHandleResponse() (*models.JumboProduct, error) {
	return jumboData, nil
}
