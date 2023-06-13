package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"webScraper/src/constants"
	"webScraper/src/interfaces"
	"webScraper/src/models/auth"
	"webScraper/src/models/scraping"
	services "webScraper/src/services/emails"

	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

func TrackProduct(product scraping.Product, userId string) error {
	// TODO: Reduce this function using sub functions
	// Open the connection now to the scraping DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error Creating the connection to the webscraping DB ", err)
		return err
	}

	// Close the connection
	defer func() {
		err = CloseConnection(connection)
		if err != nil {
			log.Println("Error closing in users ", err)
			return
		}
	}()

	// Check if the product already exists
	exists := checkIfProductExists(connection, product.Product_id)
	if exists == nil {
		log.Println("The product already exists")

		// Check if the user has tracked the product
		isTracked, err := checkIfProductIsTrackedByUser(product.Product_id, userId)
		if err != nil {
			return err
		}

		if !isTracked {
			// Add the tracking to the account
			err = createUserProductField(connection, product, userId)
			if err != nil {
				return err
			}
		}

		// Check if the price change
		priceChange, savedPrice, err := checkPriceChanges(connection, product)
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

			// Send an email if the new price is lower to the users
			if savedPrice > product.Current_price {
				// Get the UserId of the clients that are tracking the prodcut
				users, err := GetUserIdByProductId(product.Product_id)
				if err != nil {
					return err
				}

				// Prepare the data to send
				sender := gomail.NewDialer(constants.SMTP_HOST, 587, os.Getenv(constants.MY_EMAIL), os.Getenv(constants.EMAIL_PASSWORD))
				alertProduct := &scraping.Product{
					Product_id:    product.Product_id,
					ImageURL:      product.ImageURL,
					Name:          product.Name,
					Description:   product.Description,
					ProductURL:    product.ProductURL,
					Current_price: product.Current_price,
					High_price:    savedPrice, // The older and lower price
				}

				// Send the email
				for _, user := range users {
					user, err := GetUserById(user)
					if err != nil {
						log.Println("Error getting the user")
						return err
					}

					err = services.SendNotificationLowerPrice(user, sender, *alertProduct)
					if err != nil {
						log.Println("Error sending lower price email")
						return err
					}
				}
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

	return nil
}

func CreateProduct(product interfaces.Product, userId string) error {
	newProduct := product.CreateProductStructure(userId)
	err := TrackProduct(newProduct, userId)
	if err != nil {
		log.Println("Error creating the Amazon product")
		return err
	}
	return nil
}

func DeleteProduct(product interfaces.Product) {
	// Create the product structure
	generalProduct := product.CreateProductStructure("")

	// Create connection to the DB
	db, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error deleting the product, DB can't connect")
		return
	}

	// Close the connection
	defer func() {
		err = CloseConnection(db)
		if err != nil {
			log.Println("Error closing the connection in products ", err)
			return
		}
	}()

	// Delete the element from the price_history table
	sqlSentence := fmt.Sprintf("DELETE FROM price_history WHERE product_id = '%s'", generalProduct.Product_id)
	_, err = db.Exec(sqlSentence)
	if err != nil {
		log.Println("Error deleting the product in price_history table ", err)
		return
	}

	// Delete the element from the price table
	sqlSentence = fmt.Sprintf("DELETE FROM price WHERE product_id = '%s'", generalProduct.Product_id)
	_, err = db.Exec(sqlSentence)
	if err != nil {
		log.Println("Error deleting the product in price table ", err)
		return
	}

	// Delete the element from the user_product table
	sqlSentence = fmt.Sprintf("DELETE FROM user_product WHERE product_id = '%s'", generalProduct.Product_id)
	_, err = db.Exec(sqlSentence)
	if err != nil {
		log.Println("Error deleting the product in price table ", err)
		return
	}

	// Delete the element from the product table
	sqlSentence = fmt.Sprintf("DELETE FROM product WHERE product_id = '%s'", generalProduct.Product_id)
	_, err = db.Exec(sqlSentence)
	if err != nil {
		log.Println("Error deleting the product from the product table ", err)
		return
	}

	log.Printf("Product %s deleted\n", generalProduct.Product_id)
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

func checkPriceChanges(connection *sql.DB, product scraping.Product) (bool, string, error) {
	// Get the Price from the DB
	sqlSentence := fmt.Sprintf("SELECT * FROM price WHERE product_id = '%s'", product.Product_id)
	response, err := connection.Query(sqlSentence)
	if err != nil {
		log.Println("Error making the query for getting the price ", err)
		return false, "", err
	}

	// Extract the name of the product if it exits
	var price_field [5]string
	for response.Next() {
		err = response.Scan(&price_field[0], &price_field[1], &price_field[2], &price_field[3], &price_field[4])
		if err != nil {
			log.Println("Error getting the stored prices", err)
			return false, "", err
		}
	}

	// Check if the current price is equal to the previously price
	if price_field[2] == product.Current_price {
		log.Println("The price is equal")
		return false, price_field[2], nil
	}

	log.Println("The price has changed")
	defer response.Close()
	return true, price_field[2], nil
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
		log.Println("Error getting the user, DB can't connect")
		return err
	}

	// Close the connection
	defer func() {
		err = CloseConnection(connection)
		if err != nil {
			log.Println("Error closing in users ", err)
			return
		}
	}()

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
		log.Println("Error getting the user, DB can't connect")
		return false, err
	}

	// Close the connection
	defer func() {
		err = CloseConnection(connection)
		if err != nil {
			log.Println("Error closing in users ", err)
			return
		}
	}()

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

	// Close the connection
	err = response.Close()
	if err != nil {
		log.Println("Error closing the response with the DB")
		return false, err
	}
	return user_validation, nil
}

func CheckIfUserExists(userId string) (bool, error) {
	// Create connection to the DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error getting the user, DB can't connect")
		return false, err
	}

	// Close the connection
	defer func() {
		err = CloseConnection(connection)
		if err != nil {
			log.Println("Error closing in users ", err)
			return
		}
	}()

	// Check if the user exits
	sqlSentence := fmt.Sprintf("SELECT user_id FROM users WHERE user_id = '%s'", userId)
	response, err := connection.Query(sqlSentence)
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

	// Close the connection
	err = response.Close()
	if err != nil {
		log.Println("Error closing the response with the DB")
		return false, err
	}

	return true, nil
}

func GetAllProducts() ([]string, error) {
	// Create connection to the DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error getting the user, DB can't connect")
		return nil, err
	}

	// Close the connection
	defer func() {
		err = CloseConnection(connection)
		if err != nil {
			log.Println("Error closing in users ", err)
			return
		}
	}()

	// Check if the user exits
	sqlSentence := "SELECT productUrl FROM product"
	response, err := connection.Query(sqlSentence)
	if err != nil {
		log.Println("Error making the query for getting the price ", err)
		return nil, err
	}

	// Extract the name of the product if it exits
	var products []string
	var productURL string
	for response.Next() {
		err = response.Scan(&productURL)
		if err != nil {
			log.Println("User not Foud", err)
			return nil, err
		}
		products = append(products, productURL)
	}

	err = response.Close()
	if err != nil {
		log.Println("Error closing the response with the DB")
		return nil, err
	}

	return products, nil
}

func GetUserById(userId string) (*auth.User, error) {
	// Create connection to the DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error getting the user, DB can't connect")
		return nil, err
	}

	// Close the connection
	defer func() {
		err = CloseConnection(connection)
		if err != nil {
			log.Println("Error closing in users ", err)
			return
		}
	}()

	// Get the user
	sqlSentence := fmt.Sprintf("SELECT user_id,name,email FROM users WHERE user_id = '%s'", userId)
	response, err := connection.Query(sqlSentence)
	if err != nil {
		log.Println("Error making the query for getting the price ", err)
		return nil, err
	}

	user := &auth.User{}
	for response.Next() {
		err = response.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			log.Println("Error getting the user", err)
			return nil, err
		}
	}

	err = response.Close()
	if err != nil {
		log.Println("Error closing the response with the DB")
		return nil, err
	}

	return user, nil
}

func GetUserIdByProductId(productId string) ([]string, error) {
	// Create connection to the DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error getting the user, DB can't connect")
		return nil, err
	}

	// Close the connection
	defer func() {
		err = CloseConnection(connection)
		if err != nil {
			log.Println("Error closing in users ", err)
			return
		}
	}()

	// Get the user
	sqlSentence := fmt.Sprintf("SELECT user_id FROM user_product WHERE product_id = '%s'", productId)
	response, err := connection.Query(sqlSentence)
	if err != nil {
		log.Println("Error making the query for getting the price ", err)
		return nil, err
	}

	var users []string
	var userId string
	for response.Next() {
		err = response.Scan(&userId)
		if err != nil {
			log.Println("Error getting the user", err)
			return nil, err
		}
		users = append(users, userId)
	}

	err = response.Close()
	if err != nil {
		log.Println("Error closing the response with the DB")
		return nil, err
	}

	return users, nil
}

func checkIfProductIsTrackedByUser(productId string, userId string) (bool, error) {
	// Create connection to the DB
	connection, err := CreateConnectionToDatabase("webscraping")
	if err != nil {
		log.Println("Error getting the user, DB can't connect")
		return true, err
	}

	// Close the connection
	defer func() {
		err = CloseConnection(connection)
		if err != nil {
			log.Println("Error closing in users ", err)
			return
		}
	}()

	// Get the user
	sqlSentence := fmt.Sprintf("SELECT user_product_id FROM user_product WHERE user_id = '%s' AND product_id = '%s'", userId, productId)
	response, err := connection.Query(sqlSentence)
	if err != nil {
		log.Println("Error making the query for getting the price ", err)
		return true, err
	}

	var result string
	for response.Next() {
		err = response.Scan(&result)
		if err != nil {
			log.Println("Error getting the user", err)
			return true, err
		}
	}

	if result == " " {
		log.Println("Product assigned to the client account")
		return false, nil
	}

	err = response.Close()
	if err != nil {
		log.Println("Error closing the response with the DB")
		return true, err
	}

	return true, nil

}
