package main

import (
	"encoding/csv"
	"fmt"
	"log"
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
		colly.AllowedDomains("www.alkosto.com", "alkosto.com", "exito.com", "www.exito.com"),
		colly.CacheDir("./cache"),
	)
	collector.SetRequestTimeout(30 * time.Second)

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

	collector.OnHTML("script[type='application/ld+json']", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	err = collector.Visit("https://www.exito.com/televisor-samsung-55-pulgadas-uhd-4k-un55bu8000kxzl-3070371/p")
	if err != nil {
		log.Fatal("Error Visiting the page ", err)
	}
	collector.Wait()
	log.Println("Scraping Complete ")
}
