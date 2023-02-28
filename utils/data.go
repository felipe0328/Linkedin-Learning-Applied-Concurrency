package utils

import (
	"appliedConcurrency/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

const produtsInputPath = "./input/products.csv"

func ImportProductsData(products *sync.Map) error {
	data, err := readCSV(produtsInputPath)
	if err != nil {
		return err
	}

	for _, line := range data {
		// Checking the data is legitimate
		if len(line) != 5 {
			continue
		}

		id := line[0]
		stock, err := strconv.Atoi(line[2])
		if err != nil {
			continue
		}

		price, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			continue
		}

		newProduct := models.Product{
			ID:    id,
			Name:  fmt.Sprintf("%s(%s)", line[1], line[3]),
			Stock: stock,
			Price: price,
		}
		products.Store(id, newProduct)
	}

	return nil
}

func readCSV(filename string) ([][]string, error) {
	var data [][]string

	f, err := os.Open(filename)

	if err != nil {
		return data, err
	}

	defer f.Close()
	data, err = csv.NewReader(f).ReadAll()
	if err != nil {
		return data, err
	}

	return data, nil
}
