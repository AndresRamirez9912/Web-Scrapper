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

	// Create Cookie and send
	cookie := &http.Cookie{
		Name:     "First Cookie",
		Value:    "Value Of the cookie",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	// Redirect to the main page
	http.RedirectHandler("/", 200)
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

	// Create Cookie and send
	cookie := &http.Cookie{
		Name:     "First Cookie",
		Value:    "Value Of the cookie",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	// Redirect to the main page
	http.RedirectHandler("/", 200)
}
