package interfaces

import (
	"log"
	"webScraper/src/database"
	"webScraper/src/models/scraping"
)

type Product interface {
	CreateProductStructure(userId string) scraping.Product
}

func CreateProduct(product Product, userId string) error {
	newProduct := product.CreateProductStructure(userId)
	err := database.TrackProduct(newProduct, userId)
	if err != nil {
		log.Println("Error creating the Amazon product")
		return err
	}
	return nil
}
