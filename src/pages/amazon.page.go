package pages

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"webScraper/src/constants"
	"webScraper/src/models"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var amazonData string

func InitAmazonCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN),
		colly.CacheDir(constants.CACHE),
	)
	collector.SetRequestTimeout(120 * time.Second)
	extensions.RandomUserAgent(collector) // Assign a random User Agent
	return collector
}

func AmazonOnRequest(r *colly.Request) {
	log.Println("Visiting", r.URL)
}

func AmazonOnError(r *colly.Response, err error) {
	log.Fatal("Error: ", err)
}

func AmazonOnResponse(r *colly.Response) {
	log.Println("Response Code: ", r.StatusCode)
}

func AmazonOnHTML(h *colly.HTMLElement) {
	// amazonData = h.Text // Send the response
	fmt.Println("Entro!!!!!")
	fmt.Println(h.ChildText("h1"))
}

func AmazonHandleResponse() (*models.ExitoProduct, error) {
	AmazonProduct := &models.ExitoProduct{}
	err := json.Unmarshal([]byte(amazonData), AmazonProduct)
	if err != nil {
		log.Fatal("Error unmarshaling scraping response ", err)
		return nil, err
	}
	return AmazonProduct, nil
}
