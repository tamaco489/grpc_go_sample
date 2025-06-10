package entity

// Product represents the domain entity for product
type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Categories  []string
}
