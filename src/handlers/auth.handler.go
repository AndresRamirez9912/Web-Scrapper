package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"
	"webScraper/src/constants"
	"webScraper/src/database"
	"webScraper/src/models/auth"
	services "webScraper/src/services/emails"
	"webScraper/src/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Read the body
	user, err := utils.GetBody(r)
	if err != nil {
		log.Println("Error Reading the user from Body")
	}

	// Fill the user data
	cookie_session := uuid.NewString()

	user.Id = uuid.New().String()
	user.Session_cookie = cookie_session

	// Create the user into the DB
	err = database.CreateUser(user)
	if err != nil {
		log.Println("Error creating the user from Body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the Verification Email
	err = sendVerificationEmail(user)
	if err != nil {
		log.Println("Error sending the email verification")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create Cookie and send it
	cookie := &http.Cookie{
		Name:     constants.SESSION_TOKEN,
		Value:    cookie_session,
		Expires:  time.Now().Add(60 * time.Minute), // Each session expires after 1 hour
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	// Success Response
	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)

	// Redirect to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Read the body
	user, err := utils.GetBody(r)
	if err != nil {
		log.Println("Error Reading the user from Body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if the user exists
	credentials, err := database.GetUserByEmail(user.Email)
	if err != nil {
		log.Println("Email doesn't exist")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(credentials), []byte(user.Password))
	if err != nil {
		log.Println("Incorrect Password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new cookie (because the user login)
	session_cookie := uuid.NewString()
	cookie := &http.Cookie{
		Name:     constants.SESSION_TOKEN,
		Value:    session_cookie,
		Expires:  time.Now().Add(60 * time.Minute), // Each session expires after 1 hour
		HttpOnly: true,
	}

	// Update Cookie in DB
	err = database.UpdateCookie(user.Email, session_cookie)
	if err != nil {
		log.Println("Error updating the cookie ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response to the client
	http.SetCookie(w, cookie)
	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)

	// Redirect to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func sendVerificationEmail(user *auth.User) error {
	var body bytes.Buffer

	// Get the Template with the values
	template, err := template.ParseFiles("src/services/emails/templates/emailVerification.template.html")
	if err != nil {
		log.Fatal("Error Trying to get the template ", err)
		return err
	}

	// Create the struct with the data to send to the template
	data := struct {
		UserName string
		Link     string
	}{
		UserName: user.Name,
		Link:     "google.com",
	}

	// Execute the template and get the string
	err = template.Execute(&body, data)
	if err != nil {
		log.Fatal("Error Trying to execute the template ", err)
		return err
	}

	// Send the email
	err = services.SendEmail(user.Email, constants.ACCOUNT_VERIFICATION_SUBJECT, body.String())
	if err != nil {
		log.Fatal("Error Sending the email", err)
		return err
	}
	return nil
}
