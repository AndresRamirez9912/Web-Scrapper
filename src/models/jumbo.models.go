package models

type JumboProduct struct {
	Name         string
	CurrentPrice string
	Disccount    string
	HighPrice    string
	Brand        string
	Description  map[int]string
	ImageURL     string
}
