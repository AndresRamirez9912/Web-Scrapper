package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"webScraper/src/models/auth"
	models "webScraper/src/models/scraping"
)

func GetProductURL(r *http.Request) (string, error) {
	// Get the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error receiving the request ", err)
		return "", err
	}

	// Convert the data into an object
	product := &models.ProductRequest{}
	err = json.Unmarshal(body, product)
	if err != nil {
		log.Println("Error Getting the data from the body request ", err)
		return "", err
	}
	return product.ProductURL, nil
}

func GetBody(r *http.Request) (*auth.User, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading the body request")
		return nil, err
	}
	user := &auth.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		log.Println("Error unmarshalling the body request")
		return nil, err
	}
	return user, nil
}
