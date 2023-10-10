package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"webScraper/src/constants"
	"webScraper/src/database"
	"webScraper/src/interfaces"
	"webScraper/src/models/auth"
	models "webScraper/src/models/scraping"
)

func GetProductURL(r *http.Request) (string, error) {
	// Get the body
	body, err := io.ReadAll(r.Body)
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
	body, err := io.ReadAll(r.Body)
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

func StoreAndSendResponse(w http.ResponseWriter, r *http.Request, scraper interfaces.Collectors) error {
	// Get the User Id from the Cookie
	userId := r.Context().Value(constants.USER_ID).(string)

	// Create the product in the DB
	err := database.CreateProduct(scraper, userId)
	if err != nil {
		log.Println("Error creating the product in the DB")
		w.WriteHeader(http.StatusInternalServerError)
		database.DeleteProduct(scraper)
		return err
	}

	// Send the response
	dataResponse, err := json.Marshal(scraper)
	if err != nil {
		log.Println("Error Serializing the response", err)
		return err
	}

	w.Header().Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(dataResponse) // Send body
	if err != nil {
		log.Println("Error sending the response")
		w.WriteHeader(http.StatusInternalServerError)
		database.DeleteProduct(scraper)
		return err
	}
	return nil
}
