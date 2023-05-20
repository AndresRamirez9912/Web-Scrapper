package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"webScraper/src/models/scraping"

	"github.com/google/uuid"
)

func TrackProduct(product scraping.Product, userId string) error {
	// Open the connection now to the scraping DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error Creating the connection to the webscraping DB ", err)
		return err
	}

	// Check if the product already exists
	exists := checkIfProductExists(connection, product.Product_id)
	if exists == nil {
		log.Println("The product already exists")

		// Check if the price change
		priceChange, err := checkPriceChanges(connection, product)
		if err != nil {
			return errors.New("Error checking the price")
		}

		// If the price has chaged, update the current price and store in historical
		if priceChange {
			// Update Current Price
			err = updatePrice(connection, product)
			if err != nil {
				return err
			}

			// Store in Price History DB
			err = createPriceHistoryField(connection, product)
			if err != nil {
				return err
			}

		}
		return nil
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

	// Create the Price History field in the DB
	err = createPriceHistoryField(connection, product)
	if err != nil {
		return err
	}

	// Close the connection
	defer CloseConnection(connection)

	return nil
}

func checkIfProductExists(connection *sql.DB, product_id string) error {
	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("SELECT product_id FROM product WHERE product_id = '%s'", product_id)
	response, err := connection.Query(sqlSentence)
	if err != nil {
		return err
	}

	// Extract the name of the product if it exits
	var product string
	for response.Next() {
		err = response.Scan(&product)
		if err != nil {
			return err
		}
	}

	if product == "" {
		defer response.Close()
		return errors.New("The product doesn't exists")
	}

	defer response.Close()
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

func checkPriceChanges(connection *sql.DB, product scraping.Product) (bool, error) {
	// Get the Price from the DB
	sqlSentence := fmt.Sprintf("SELECT * FROM price WHERE product_id = '%s'", product.Product_id)
	response, err := connection.Query(sqlSentence)
	if err != nil {
		log.Println("Error making the query for getting the price ", err)
		return false, err
	}

	// Extract the name of the product if it exits
	var price_field [5]string
	for response.Next() {
		err = response.Scan(&price_field[0], &price_field[1], &price_field[2], &price_field[3], &price_field[4])
		if err != nil {
			log.Println("Error getting the stored prices", err)
			return false, err
		}
	}

	// Check if the current price is equal to the previously price
	if price_field[2] == product.Current_price {
		log.Println("The price is equal")
		return false, nil
	}

	log.Println("The price has changed")
	defer response.Close()
	return true, nil
}

func updatePrice(connection *sql.DB, product scraping.Product) error {
	sentence := `UPDATE price SET current_price='%s', discount='%s', high_price='%s' WHERE product_id='%s'`

	sqlSentence := fmt.Sprintf(sentence, product.Current_price, product.Discount, product.High_price, product.Product_id)

	// Execute the SQL command
	_, err := connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Error Updating the Price field", err)
		return err
	}

	log.Println("Price Updated Successfully")
	return nil
}

func VerifyEmail(userId string) error {
	// Create connection to the DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Fatal("Error getting the user, DB can't connect")
		return err
	}

	sentence := `UPDATE users SET email_verification='TRUE' WHERE user_id='%s'`
	sqlSentence := fmt.Sprintf(sentence, userId)

	// Execute the SQL command
	_, err = connection.Exec(sqlSentence)
	if err != nil {
		log.Println("Errorupdating the email validation field", err)
		return err
	}

	log.Println("Email validated successfully Successfully")
	return nil
}

func CheckUserValidated(userId string) (bool, error) {
	// Create connection to the DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Fatal("Error getting the user, DB can't connect")
		return false, err
	}

	// Execute the SQL sentence
	sqlSentence := fmt.Sprintf("SELECT email_verification FROM users WHERE user_id = '%s'", userId)
	response, err := connection.Query(sqlSentence)
	if err != nil {
		return false, err
	}

	// Extract the name of the product if it exits
	var user_validation bool
	for response.Next() {
		err = response.Scan(&user_validation)
		if err != nil {
			return false, err
		}
	}

	defer response.Close()
	return user_validation, nil
}

func CheckIfUserExists(userId string) (bool, error) {
	// Create connection to the DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Fatal("Error getting the user, DB can't connect")
		return false, err
	}

	// Check if the user exits
	sqlSentence := fmt.Sprintf("SELECT user_id FROM users WHERE user_id = '%s'", userId)
	response, err := connection.Query(sqlSentence)
	defer response.Close()
	if err != nil {
		log.Println("Error making the query for getting the price ", err)
		return false, err
	}

	// Extract the name of the product if it exits
	var user string
	for response.Next() {
		err = response.Scan(&user)
		if err != nil {
			log.Println("User not Foud", err)
			return false, err
		}
	}

	// Check if the current price is equal to the previously price
	if user == userId {
		return true, nil
	}

	return true, nil
}
