package interfaces

import (
	"webScraper/src/models/scraping"
)

type Product interface {
	CreateProductStructure(userId string) scraping.Product
}
