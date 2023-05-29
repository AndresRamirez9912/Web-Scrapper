package interfaces

import "github.com/gocolly/colly"

type Collectors interface {
	OnError(colly.ErrorCallback)
	OnRequest(colly.RequestCallback)
	OnResponse(colly.ResponseCallback)
	OnHTML(string, colly.HTMLCallback)
	Visit(string) error
	Wait()
}
