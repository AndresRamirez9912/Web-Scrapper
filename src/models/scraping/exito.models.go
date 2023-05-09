package scraping

import "strconv"

type ExitoProduct struct {
	Id          string
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image"`
	Discount    string
	ProductURL  string `json:"@id"`
	Offers      offers `json:"offers"`
	Brand       struct {
		Name string `json:"name"`
	} `json:"brand"`
}

type offers struct {
	LowPrice  int    `json:"lowPrice"`
	HighPrice int    `json:"highPrice"`
	Currency  string `json:"priceCurrency"`
}

func (E *ExitoProduct) CreateProductStructure(userId string) Product {
	current := strconv.Itoa(E.Offers.LowPrice)
	high := strconv.Itoa(E.Offers.HighPrice)
	return Product{
		Product_id:      E.Id,
		User_product_id: userId + E.Id,
		Name:            E.Name,
		Brand:           E.Brand.Name,
		Description:     E.Description,
		ImageURL:        E.ImageURL,
		ProductURL:      E.ProductURL,
		Current_price:   current,
		Discount:        E.Discount,
		High_price:      high,
	}
}
