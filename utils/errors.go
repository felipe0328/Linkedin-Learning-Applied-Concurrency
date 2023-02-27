package utils

import "fmt"

var (
	OrderNotFound = "no order found for %s order id"
	ProductNotFound = "no product found for id %s"
)

func Error(errorMessage string, data ...any) error {
	return fmt.Errorf(errorMessage, data...)
}
