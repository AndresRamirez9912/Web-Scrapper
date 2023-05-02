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
		log.Fatal("Error hashing the password")
		return err
	}

	// Create connection to the DB
	db, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Fatal("Error creating the user, DB can't connect")
		return err
	}

	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("INSERT INTO users (id, name, email, password) VALUES ('%s', '%s', '%s', '%s');", user.Id, user.Name, user.Email, string(hash))
	_, err = db.Exec(sqlSentence)
	if err != nil {
		log.Fatal("Error Creating the the user ", err)
		return err
	}

	// Close the connection
	err = CloseConnection(db)
	if err != nil {
		log.Fatal("Error closing in users ", err)
		return err
	}
	return nil
}

func GetUserByEmail(email string) (string, error) {
	// Create connection to the DB
	db, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Fatal("Error creating the user, DB can't connect")
		return "", err
	}

	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("SELECT password FROM users WHERE email = '%s'", email)
	response, err := db.Query(sqlSentence)
	if err != nil {
		log.Fatal("Error Creating the the user ", err)
		return "", err
	}

	// Extract the password of the result
	var password string
	for response.Next() {
		err = response.Scan(&password)
		if err != nil {
			log.Println("Error loading the password")
		}
	}
	defer response.Close()
	return password, nil
}
