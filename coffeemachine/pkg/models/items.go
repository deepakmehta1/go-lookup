package models

// Item represents a product in the vending machine.
type Item struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}
