package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"webScraper/src/models"
)

func GetProductURL(r *http.Request) (string, error) {
	// Get the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error receiving the request ", err)
		return "", err
	}

	// Convert the data into an object
	product := &models.ProductRequest{}
	err = json.Unmarshal(body, product)
	if err != nil {
		log.Fatal("Error Getting the data from the body request ", err)
		return "", err
	}
	return product.ProductURL, nil
}
