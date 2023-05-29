package scrapers

import (
	"fmt"
	"testing"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
)

func TestSendAmazonCollyRequest(t *testing.T) {
	t.Run("Success Flow", func(t *testing.T) {
		// Create Object interface
		scraper := interfaces.Scraper{
			AllowedDomains: []string{constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN},
		}

		// Create Collector
		collector := scraper.InitCollector()
		fmt.Println(collector)
	})
}
