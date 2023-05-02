package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func CreateConnectionToDatabase(database string) (*sql.DB, error) {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbName := database
	psqlData := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	if database == "" {
		psqlData = strings.Replace(psqlData, "dbname=", "", -1)
	}

	// Open connection
	db, err := sql.Open("postgres", psqlData)
	if err != nil {
		log.Println("Error creating the connection to the dataBase ", err)
		return nil, err
	}

	// Check Connection
	if err = db.Ping(); err != nil {
		log.Println("Error trying to connect to the dataBase ", err)
		return nil, err
	}
	return db, nil
}

func CloseConnection(db *sql.DB) error {
	return db.Close()
}

func CreateTables(db *sql.DB) {
	// Create SQL Database
	sqlSentence := "CREATE DATABASE webscraping"
	_, err := db.Exec(sqlSentence)
	if err != nil {
		log.Fatal("Error Creating the DataBase Web Scraping", err)
	}

	// Close the general connection
	err = CloseConnection(db)
	if err != nil {
		log.Fatal("Error Closing the connection to the general DB")
	}

	// Open the connection now to the scraping DB
	dbScraping, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error Creating the connection to the webscraping DB ", err)
	}

	//  Create SQL Tables
	sqlSentence = `
	CREATE TABLE users (
		id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}$'),
		password VARCHAR(100) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	// Execute the SQL command
	_, err = dbScraping.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the Table Users", err)
	}

	log.Println("Successfully creation of the users table")

	// Close the connection
	defer CloseConnection(dbScraping)
}

func ClearDatabase() error {
	// Open the connection now to the scraping DB
	db, err := CreateConnectionToDatabase("")
	if err != nil {
		log.Println("Error Creating the connecton to DB ", err)
		return err
	}

	sqlSentence := "DROP DATABASE webScraping"
	_, err = db.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the DataBase Web Scraping", err)
		return err
	}

	// Close connection
	CloseConnection(db)
	if err != nil {
		log.Println("Error Closing the connecton to the webscraping DB ", err)
		return err
	}
	return nil
}

func ChecWebScrapingExist() bool {
	// Create connection
	dbScraping, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Creating the webscraping database")
		return false
	}

	// Close connection
	err = CloseConnection(dbScraping)
	if err != nil {
		log.Println("Error Closing the connecton to the webscraping DB ", err)
		return false
	}
	log.Println("Database webscraping already exists")
	return true // The DB exists
}
