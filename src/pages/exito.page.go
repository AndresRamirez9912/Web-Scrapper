package pages

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"webScraper/src/models"

	"github.com/gocolly/colly"
)

func InitExitoCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("exito.com", "www.exito.com"),
		colly.CacheDir("./cache"),
	)
	collector.SetRequestTimeout(120 * time.Second)
	return collector
}

func ExitoOnRequest(r *colly.Request) {
	r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	fmt.Println("Visiting", r.URL)
}

func ExitoOnError(r *colly.Response, err error) {
	log.Fatal("Error: ", err)
}

func ExitoOnResponse(r *colly.Response) {
	fmt.Println("Response Code: ", r.StatusCode)
}

func ExitoOnHTML(h *colly.HTMLElement) {
	// fmt.Println(h.Text)
	handleResponse(h.Text)
}

func handleResponse(data string) {
	exitoData := &models.ExitoProduct{}
	err := json.Unmarshal([]byte(data), exitoData)
	if err != nil {

	}
	fmt.Printf("%+v\n\n", exitoData)
}
