package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"webScraper/src/models/scraping"
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

	// Create relation to the user inserting a field in the user_product table
	err = createUserProductField(connection, product, userId)
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
		log.Println("Error Creating the user Table", err)
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
		log.Println("Error Creating the user Table", err)
		return err
	}

	log.Println("User Created Successfully")
	return nil
}
