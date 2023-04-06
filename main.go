package main

import (
	"log"
	"webScraper/src/pages"
)

func main() {

	// Create a collector to setup the data searcher
	collector := pages.InitExitoCollector()

	// Callbacks
	collector.OnError(pages.ExitoOnError)

	collector.OnRequest(pages.ExitoOnRequest)

	collector.OnResponse(pages.ExitoOnResponse)

	collector.OnHTML("script[type='application/ld+json']", pages.ExitoOnHTML)

	// Visit the page
	err := collector.Visit("https://www.exito.com/leche-descremada-sixpack-x-1300-ml-cu-563046/p")
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}

	collector.Wait()
}
