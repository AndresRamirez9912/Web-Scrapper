package handlers

import (
	"log"
	"net/http"
	"webScraper/src/database"
	"webScraper/src/utils"

	"golang.org/x/crypto/bcrypt"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	// Read the body
	user, err := utils.GetBody(r)
	if err != nil {
		log.Println("Error Reading the user from Body")
	}

	// Create the user
	err = database.CreateUser(user)
	if err != nil {
		log.Println("Error creating the user from Body")
	}

	// Success Response
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(201)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Read the body
	user, err := utils.GetBody(r)
	if err != nil {
		log.Println("Error Reading the user from Body")
	}

	// Compare the credentials
	credentials, err := database.GetUserByEmail(user.Email)
	if err != nil {
		log.Println("Email doesn't exist")
	}
	err = bcrypt.CompareHashAndPassword([]byte(credentials), []byte(user.Password))
	if err != nil {
		log.Println("Incorrect Password")
		w.WriteHeader(401)
		return
	}

	// Response
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
}
