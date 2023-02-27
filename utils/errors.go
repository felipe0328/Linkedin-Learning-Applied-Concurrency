package utils

import "fmt"

var (
	// Orders
	OrderNotFound    = "no order found for %s order id"
	OrderAmountWrong = "order amount must be at least 1:got %d"

	// Products
	ProductNotFound  = "no product found for id %s"
	ProductNotExists = "product %s does not exists"
	ProductNotStock = "not enough stock for product %s:got %d, want %d"
)

func Error(errorMessage string, data ...any) error {
	return fmt.Errorf(errorMessage, data...)
}
