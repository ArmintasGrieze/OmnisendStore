package models

type Product struct {
	ID          int
	Title       string
	Description string
	Quantity    int
	Model       string
	Year        int
	Price       float64
	ImageURL    string
}
