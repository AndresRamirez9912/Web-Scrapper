package models

type AmazonProduct struct {
	Name        string
	LowPrice    string
	Disccount   string
	HighPrice   string
	Brand       string
	Description map[int]string
	ImageURL    string
}
