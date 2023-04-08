package pages

import (
	"encoding/json"
	"log"
	"time"
	"webScraper/src/constants"
	"webScraper/src/models"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var exitoData string

func InitExitoCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(constants.EXITO_HALF_DOMAIN, constants.EXITO_DOMAIN),
		colly.CacheDir(constants.CACHE),
	)
	collector.SetRequestTimeout(120 * time.Second)
	extensions.RandomUserAgent(collector) // Assign a random User Agent
	return collector
}

func ExitoOnRequest(r *colly.Request) {
	log.Println("Visiting", r.URL)
}

func ExitoOnError(r *colly.Response, err error) {
	log.Fatal("Error: ", err)
}

func ExitoOnResponse(r *colly.Response) {
	log.Println("Response Code: ", r.StatusCode)
}

func ExitoOnHTML(h *colly.HTMLElement) {
	exitoData = h.Text // Send the response
}

func ExitoHandleResponse() (*models.ExitoProduct, error) {
	exitoProduct := &models.ExitoProduct{}
	err := json.Unmarshal([]byte(exitoData), exitoProduct)
	if err != nil {
		log.Fatal("Error unmarshaling scraping response ", err)
		return nil, err
	}
	return exitoProduct, nil
}
