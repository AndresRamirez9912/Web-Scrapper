package scraping

type JumboProduct struct {
	Id           string
	Name         string
	CurrentPrice string
	Disccount    string
	HighPrice    string
	ProductURL   string
	Brand        string
	Description  map[int]string
	ImageURL     string
}

func (A *JumboProduct) CreateProductStructure(userId string) Product {
	description := "&"
	for _, v := range A.Description {
		description = v + description
	}
	return Product{
		Product_id:      A.Id,
		User_product_id: userId + A.Id,
		Name:            A.Name,
		Brand:           A.Brand,
		Description:     description,
		ImageURL:        A.ImageURL,
		ProductURL:      A.ProductURL,
		Current_price:   A.CurrentPrice,
		Discount:        A.Disccount,
		High_price:      A.HighPrice,
	}
}
