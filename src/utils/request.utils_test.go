package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"webScraper/src/models/auth"
	models "webScraper/src/models/scraping"
)

func TestGetProductURL(t *testing.T) {

	t.Run("Success Result", func(t *testing.T) {
		desiredURL := "https://www.exito.com/scooter-electrica-vsett-9-19ah-102553632-mp/p"

		bodyRequest := &models.ProductRequest{
			ProductURL: "https://www.exito.com/scooter-electrica-vsett-9-19ah-102553632-mp/p",
		}
		jsonBody, err := json.Marshal(bodyRequest)
		if err != nil {
			t.Error("Error ", err)
		}

		request, err := http.NewRequest("", "", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Error("Error creting the request ", err)
		}

		productURL, err := GetProductURL(request)
		if err != nil {
			t.Error("Error getting the URL ", err)
		}

		if productURL != desiredURL {
			t.Error("The desired URL didn't match with the received URL")
		}
	})

	t.Run("Fail Parsing the data into the struct", func(t *testing.T) {
		request, err := http.NewRequest("", "", bytes.NewBuffer([]byte("https://www.exito.com/scooter-electrica-vsett-9-19ah-102553632-mp/p")))
		if err != nil {
			t.Error("Error creting the request ", err)
		}

		_, err = GetProductURL(request)
		if err == nil {
			t.Error("The code must throw an error ")
		}
	})

	t.Run("Fail Getting the body", func(t *testing.T) {

		request, err := http.NewRequest("", "", errReader(0))
		if err != nil {
			t.Error("Error creting the request ", err)
		}

		_, err = GetProductURL(request)
		if err == nil {
			t.Error("The code must throw an error ")
		}
	})
}

func TestGetBody(t *testing.T) {
	t.Run("Successful result", func(t *testing.T) {
		bodyRequest := &auth.User{
			Id:       "",
			Name:     "Andres",
			Password: "123456789",
			Email:    "andres@gmail.com",
		}
		jsonBody, err := json.Marshal(bodyRequest)
		if err != nil {
			t.Error("Error ", err)
		}

		request, err := http.NewRequest("", "", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Error("Error creting the request ", err)
		}

		user, err := GetBody(request)
		if err != nil {
			t.Error("Error getting the URL ", err)
		}

		if user.Name != bodyRequest.Name || user.Email != bodyRequest.Email {
			t.Error("The user didn't match with the body request")
		}
	})

	t.Run("Fail Getting the body", func(t *testing.T) {

		request, err := http.NewRequest("", "", errReader(0))
		if err != nil {
			t.Error("Error creting the request ", err)
		}

		_, err = GetBody(request)
		if err == nil {
			t.Error("The code must throw an error ")
		}
	})

	t.Run("Fail Parsing the user data into the struct", func(t *testing.T) {
		request, err := http.NewRequest("", "", bytes.NewBuffer([]byte("https://www.exito.com/scooter-electrica-vsett-9-19ah-102553632-mp/p")))

		if err != nil {
			t.Error("Error creting the request ", err)
		}

		_, err = GetBody(request)
		if err == nil {
			t.Error("The code must throw an error, the body json didn't match with the user struct")
		}
	})
}
