package scrapers

import (
	"testing"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/mocks"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestSendAmazonCollyRequest(t *testing.T) {
	t.Run("Success Flow", func(t *testing.T) {
		// Create Object interface
		scraper := interfaces.Scraper{
			AllowedDomains: []string{constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN},
		}

		// Create the Mock Collector
		mockCollector := mocks.CollectorMock{}

		// Execute the function
		product, err := SendAmazonCollyRequest("https://www.amazon.com/-/es/DJI-Gafas-Integra-port%C3%A1tiles-transmisi%C3%B3n/dp/B0BQYQL9C2/ref=pd_ci_mcx_mh_mcx_views_0", scraper, mockCollector)
		if err != nil {
			t.Error("An error occrred sending the colly request")
		}

		// Check the result
		if product.ProductURL != "https://www.amazon.com/-/es/DJI-Gafas-Integra-port%C3%A1tiles-transmisi%C3%B3n/dp/B0BQYQL9C2/ref=pd_ci_mcx_mh_mcx_views_0" {
			t.Error("The desired ProductURL didn't match with the received")
		}
		if product.Id != "B0BQYQL9C2" {
			t.Error("The desired Id didn't match with the received")
		}
	})

	t.Run("Fail on visit function", func(t *testing.T) {
		// Create Object interface
		scraper := interfaces.Scraper{
			AllowedDomains: []string{constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN},
		}

		// Create the Mock Collector
		mockCollector := mocks.CollectorMock{
			Fail: true,
		}

		// Execute the function and validate the error
		_, err := SendAmazonCollyRequest("https://www.amazon.com/-/es/DJI-Gafas-Integra-port%C3%A1tiles-transmisi%C3%B3n/dp/B0BQYQL9C2/ref=pd_ci_mcx_mh_mcx_views_0", scraper, mockCollector)
		if err == nil {
			t.Error("The function must return an error")
		}
	})

	t.Run("Fail getting the Id of the product", func(t *testing.T) {
		// Create Object interface
		scraper := interfaces.Scraper{
			AllowedDomains: []string{constants.AMAZON_HALF_DOMAIN, constants.AMAZON_DOMAIN},
		}

		// Create the Mock Collector
		mockCollector := mocks.CollectorMock{}

		// Execute the function and validate the error
		_, err := SendAmazonCollyRequest("https://www.amazon.com", scraper, mockCollector)
		if err == nil {
			t.Error("The function must return an error")
		}
	})
}

func TestAmazonOnHTML(t *testing.T) {
	// TODO: Create a complete Mock with an Amazon response
	data := colly.HTMLElement{
		DOM: &goquery.Selection{
			Nodes: []*html.Node{{
				Parent:      &html.Node{},
				FirstChild:  &html.Node{},
				LastChild:   &html.Node{},
				PrevSibling: &html.Node{},
				NextSibling: &html.Node{},
				Type:        html.ElementNode,
				DataAtom:    atom.Li,
				Data:        "li",
				Namespace:   "",
				Attr: []html.Attribute{
					{
						Namespace: "",
						Key:       "class",
						Val:       "a-unordered-list a-nostyle a-horizontal list maintain-height",
					},
				},
			}},
		},
	}
	amazonOnHTML(&data)
	// Check if the information was separated correctly
}
