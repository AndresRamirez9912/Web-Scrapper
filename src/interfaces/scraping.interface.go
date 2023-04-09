package interfaces

import (
	"log"
	"time"
	"webScraper/src/constants"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Scraper struct {
	AllowedDomains []string
}

type scraperMethods interface {
	InitCollector() *colly.Collector
	OnRequest(r *colly.Request)
	OnResponse(r *colly.Response)
	OnError(r *colly.Response, err error)
}

func (s Scraper) InitCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(s.AllowedDomains...),
		colly.CacheDir(constants.CACHE),
	)
	collector.SetRequestTimeout(120 * time.Second)
	extensions.RandomUserAgent(collector) // Assign a random User Agent

	return collector
}

func (s Scraper) OnRequest(r *colly.Request) {
	log.Println("Visiting", r.URL)
}

func (s Scraper) OnResponse(r *colly.Response) {
	log.Println("Response Code: ", r.StatusCode)
}

func (s Scraper) OnError(r *colly.Response, err error) {
	log.Fatal("Error: ", err)
}