package pages

import (
	"log"
	"time"
	"webScraper/src/constants"
	"webScraper/src/models"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
)

var amazonData = &models.AmazonProduct{}

func InitAmazonCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN),
		colly.CacheDir(constants.CACHE),
	)
	collector.SetRequestTimeout(120 * time.Second)
	extensions.RandomUserAgent(collector) // Assign a random User Agent

	// Set Proxy
	proxySwitcher, err := proxy.RoundRobinProxySwitcher(
		"socks5://188.226.141.127:1080",
		"socks5://67.205.132.241:1080",
		"http://103.155.62.173:8080",
	)

	if err != nil {
		log.Fatal(err)
	}
	collector.SetProxyFunc(proxySwitcher)
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
	log.Println("Proxy Usado: ", r.Request.ProxyURL)
}

func AmazonOnHTML(h *colly.HTMLElement) {
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
	amazonData.LowPrice = h.ChildText("div.a-section.a-spacing-none.aok-align-center span.a-offscreen")                                     // Product Lower Price
	amazonData.HighPrice = h.ChildText("div.a-section.a-spacing-small.aok-align-center span.a-offscreen")                                   // Original Price, withou Discount

}

func AmazonHandleResponse() (*models.AmazonProduct, error) {
	return amazonData, nil
}
