package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"webScraper/src/models/scraping"

	"github.com/google/uuid"
)

func CreateProduct(product scraping.Product, userId string) error {
	// Open the connection now to the scraping DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error Creating the connection to the webscraping DB ", err)
		return err
	}

	// Check if the product already exists
	exists := checkProduct(connection, product.Product_id)
	if exists == nil {
		return errors.New("The product already exists")
	}

	// Create the product in the DB Product
	err = createProductField(connection, product)
	if err != nil {
		return err
	}

	// Create the user_product in the DB Product
	err = createUserProductField(connection, product, userId)
	if err != nil {
		return err
	}

	// Create the price in the DB Product
	err = createPriceField(connection, product)
	if err != nil {
		return err
	}

	// Close the connection
	defer CloseConnection(connection)

	return nil
}

func checkProduct(connection *sql.DB, product_id string) error {
	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("SELECT name FROM product WHERE product_id = '%s'", product_id)
	response, err := connection.Query(sqlSentence)
	if err != nil {
		return err
	}

	// Extract the name of the product if it exits
	var name string
	for response.Next() {
		err = response.Scan(&name)
		if err != nil {
			return err
		}
	}

	if name == "" {
		defer response.Close()
		return errors.New("The product doesn't exists")
	}

	defer response.Close()
	log.Println("The product already exists")
	return nil
}

func createProductField(connection *sql.DB, product scraping.Product) error {
	sentence := `INSERT INTO product (product_id, 
		name, brand, description, imageURL, productURL) 
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s')`

	sqlSentence := fmt.Sprintf(sentence, product.Product_id,
		product.Name, product.Brand, product.Description,
		product.ImageURL, product.ProductURL)

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the product field", err)
		return err
	}

	log.Println("Product Created Successfully")
	return nil
}

func createUserProductField(connection *sql.DB, product scraping.Product, userId string) error {
	sentence := `INSERT INTO user_product (user_product_id, user_id, product_id) 
		VALUES ('%s', '%s', '%s')`

	sqlSentence := fmt.Sprintf(sentence, product.User_product_id, userId, product.Product_id)

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the user_product field", err)
		return err
	}

	log.Println("UserProduct Created Successfully")
	return nil
}

func createPriceField(connection *sql.DB, product scraping.Product) error {
	sentence := `INSERT INTO price (price_id, product_id, 
		current_price, discount, high_price) 
		VALUES ('%s', '%s', '%s', '%s', '%s')`

	sqlSentence := fmt.Sprintf(sentence, uuid.New().String(),
		product.Product_id, product.Current_price, product.Discount,
		product.High_price)

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the price field", err)
		return err
	}

	log.Println("Price Created Successfully")
	return nil
}

func createPriceHistoryField(connection *sql.DB, product scraping.Product) error {
	sentence := `INSERT INTO price_history (price_history_id, product_id, 
		current_price, discount, high_price) 
		VALUES ('%s', '%s', '%s', '%s', '%s')`

	sqlSentence := fmt.Sprintf(sentence, uuid.New().String(),
		product.Product_id, product.Current_price, product.Discount,
		product.High_price)

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Creating the Price History field", err)
		return err
	}

	log.Println("Price History Created Successfully")
	return nil
}
