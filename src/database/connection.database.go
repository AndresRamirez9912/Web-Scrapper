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

func CreateTables(db *sql.DB) error {
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
		return err
	}

	// Open the connection now to the scraping DB
	dbScraping, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error Creating the connection to the webscraping DB ", err)
		return err
	}

	// Createthe tables
	err = createUserTable(dbScraping)
	if err != nil {
		return err
	}

	err = createUserProductTable(dbScraping)
	if err != nil {
		return err
	}

	err = createProductTable(dbScraping)
	if err != nil {
		return err
	}

	err = createPriceHistoryTable(dbScraping)
	if err != nil {
		return err
	}

	err = createPriceTable(dbScraping)
	if err != nil {
		return err
	}

	err = addForeignKeys(dbScraping)
	if err != nil {
		return err
	}

	// Close the connection
	defer CloseConnection(dbScraping)
	return nil
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

func createUserTable(connection *sql.DB) error {
	sqlSentence := `
	CREATE TABLE users (
		user_id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}$'),
		password VARCHAR(100) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		session_cookie VARCHAR(50) NOT NULL UNIQUE
	)`

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the user Table", err)
		return err
	}

	log.Println("Successfully creation of the user table")
	return nil
}

func createProductTable(connection *sql.DB) error {
	sqlSentence := `
	CREATE TABLE product (
		product_id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(100) NOT NULL,		
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		brand VARCHAR(50) NOT NULL,
		description VARCHAR NOT NULL,
		imageURL VARCHAR(500) NOT NULL,
		productURL VARCHAR(500) NOT NULL
	)`

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the Product Table", err)
		return err
	}

	log.Println("Successfully creation of the Product table")
	return nil
}

func createUserProductTable(connection *sql.DB) error {
	sqlSentence := `
	CREATE TABLE user_product (
		user_product_id VARCHAR(50) PRIMARY KEY,
		user_id VARCHAR(50) NOT NULL,
		product_id VARCHAR(50) NOT NULL
	)`

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the User_Product Table", err)
		return err
	}

	log.Println("Successfully creation of the User_Product table")
	return nil
}

func createPriceTable(connection *sql.DB) error {
	sqlSentence := `
	CREATE TABLE price (
		price_id VARCHAR(50) PRIMARY KEY,
		product_id VARCHAR(50) NOT NULL,
		current_price VARCHAR(50) NOT NULL,
		discount VARCHAR(50) NOT NULL,
		high_price VARCHAR(50) NOT NULL
	)`

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the Price Table", err)
		return err
	}

	log.Println("Successfully creation of the Price table")
	return nil
}

func createPriceHistoryTable(connection *sql.DB) error {
	sqlSentence := `
	CREATE TABLE price_history (
		price_history_id VARCHAR(50) PRIMARY KEY,
		product_id VARCHAR(50) NOT NULL,
		date_changed TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		current_price VARCHAR(50) NOT NULL,
		discount VARCHAR(50) NOT NULL,
		high_price VARCHAR(50) NOT NULL
	)`

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the Price History Table", err)
		return err
	}

	log.Println("Successfully creation of the Price History Table")
	return nil
}

func addForeignKeys(connection *sql.DB) error {
	// Add FK of the user_product Table
	sqlSentence := `
	ALTER TABLE user_product 
	ADD FOREIGN KEY (product_id) REFERENCES product(product_id),
	ADD FOREIGN KEY (user_id) REFERENCES users(user_id);

	ALTER TABLE price 
	ADD FOREIGN KEY (product_id) REFERENCES product(product_id);

	ALTER TABLE price_history 
	ADD FOREIGN KEY (product_id) REFERENCES product(product_id);
	`
	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the Foreign Keys", err)
		return err
	}

	log.Println("Successfully creation of the Foreign Keys")
	return nil
}
