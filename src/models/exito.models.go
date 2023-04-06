package models

type ExitoProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image"`
	ProductURL  string `json:"@id"`
	Offers      offers `json:"offers"`
	Brand       struct {
		Name string `json:"name"`
	} `json:"brand"`
}

type offers struct {
	LowPrice  string `json:"lowPrice"`
	HighPrice string `json:"highPrice"`
	Currency  string `json:"priceCurrency"`
}
