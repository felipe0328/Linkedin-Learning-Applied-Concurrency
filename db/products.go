package db

import (
	"appliedConcurrency/models"
	"appliedConcurrency/utils"
	"sort"
)

type IProductsDB interface {
	Exists(id string) bool
	Find(id string) (models.Product, error)
	Upsert(product models.Product)
	FindAll() []models.Product
}

type ProductsDB struct {
	products map[string]models.Product
}

func NewProducts() (IProductsDB, error) {
	p := &ProductsDB{
		products: make(map[string]models.Product),
	}

	if err := utils.ImportProductsData(p.products); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *ProductsDB) Exists(id string) bool {
	_, ok := p.products[id]
	return ok
}

func (p *ProductsDB) Find(id string) (models.Product, error) {
	prod, ok := p.products[id]
	if !ok {
		return prod, utils.Error(utils.ProductNotFound, id)
	}

	return prod, nil
}

func (p *ProductsDB) Upsert(product models.Product) {
	p.products[product.ID] = product
}

func (p *ProductsDB) FindAll() []models.Product {
	result := make([]models.Product, len(p.products))

	listCounter := 0
	for _, value := range p.products {
		result[listCounter] = value
		listCounter++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}
