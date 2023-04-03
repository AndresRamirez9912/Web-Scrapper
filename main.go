package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// Create FIle to store the data
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatal("Error creating the file ", err)
	}
	defer file.Close() // When I finish to use the file, close it

	// Write information into my file
	writter := csv.NewWriter(file)
	defer writter.Flush() // Clean the buffer

	// Create a collector to setup the data searcher
	collector := colly.NewCollector(
		colly.AllowedDomains("www.alkosto.com"),
	)
	collector.WithTransport(&http.Transport{
		DialContext:           (&net.Dialer{Timeout: 30 * time.Second}).DialContext,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 2 * time.Second,
	})

	// Callbacks
	collector.OnError(func(r *colly.Response, err error) {
		log.Fatal("Error: ", err)
	})

	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code: ", r.StatusCode)
	})

	err = collector.Visit("https://www.alkosto.com")
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}
	log.Println("Scraping Complete ")
}
