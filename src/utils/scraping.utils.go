package utils

import (
	"time"
	"webScraper/src/constants"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func InitCollector(AllowedDomains []string) *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(AllowedDomains...),
		colly.CacheDir(constants.CACHE),
	)
	collector.SetRequestTimeout(90 * time.Second)
	extensions.RandomUserAgent(collector) // Assign a random User Agent

	// Assign the Proxy
	// rp, err := proxy.RoundRobinProxySwitcher("http://20.206.106.192:8123", "http://65.109.228.231:8080")
	// if err != nil {
	// 	log.Println(err)
	// }
	// collector.SetProxyFunc(rp)
	return collector
}
