package database

import (
	"fmt"
	"log"
	"webScraper/src/models/auth"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *auth.User) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing the password")
		return err
	}

	// Create connection to the DB
	db, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error creating the user, DB can't connect")
		return err
	}

	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("INSERT INTO users (user_id, name, email, password, session_cookie) VALUES ('%s', '%s', '%s', '%s', '%s');", user.Id, user.Name, user.Email, string(hash), user.Session_cookie)
	_, err = db.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the user ", err)
		return err
	}

	// Close the connection
	err = CloseConnection(db)
	if err != nil {
		log.Println("Error closing in users ", err)
		return err
	}
	return nil
}

func GetUserByEmail(email string) (string, error) {
	// Create connection to the DB
	db, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error creating the user, DB can't connect")
		return "", err
	}

	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("SELECT password FROM users WHERE email = '%s'", email)
	response, err := db.Query(sqlSentence)
	if err != nil {
		log.Println("Error Creating the the user ", err)
		return "", err
	}

	// Extract the password of the result
	var password string
	for response.Next() {
		err = response.Scan(&password)
		if err != nil {
			log.Println("Error loading the password")
			return "", err
		}
	}
	defer response.Close()
	return password, nil
}

func GetUserbyCookie(session string) (string, error) {
	// Create connection to the DB
	db, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error getting the user, DB can't connect")
		return "", err
	}

	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("SELECT user_id FROM users WHERE session_cookie = '%s'", session)
	response, err := db.Query(sqlSentence)
	if err != nil {
		log.Println("Cookie not found ", err)
		return "", err
	}

	// Extract the id of the result
	var id string
	for response.Next() {
		err = response.Scan(&id)
		if err != nil {
			log.Println("Error loading the id")
			return "", err
		}
	}
	defer response.Close()

	return id, nil
}

func UpdateCookie(userEmail string, newCookie string) error {
	// Create connection to the DB
	db, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error getting the user, DB can't connect")
		return err
	}

	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("UPDATE users SET session_cookie='%s' WHERE email='%s'", newCookie, userEmail)
	_, err = db.Exec(sqlSentence)
	if err != nil {
		log.Println("Error updating cookie ", err)
		return err
	}

	return nil
}
