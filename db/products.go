package db

import (
	"appliedConcurrency/models"
	"appliedConcurrency/utils"
	"sort"
	"sync"
)

type IProductsDB interface {
	Exists(id string) bool
	Find(id string) (models.Product, error)
	Upsert(product models.Product)
	FindAll() []models.Product
}

type ProductsDB struct {
	products sync.Map
}

func NewProducts() (IProductsDB, error) {
	p := &ProductsDB{}

	if err := utils.ImportProductsData(&p.products); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *ProductsDB) Exists(id string) bool {
	_, ok := p.products.Load(id)
	return ok
}

func (p *ProductsDB) Find(id string) (models.Product, error) {
	var prod models.Product
	result, ok := p.products.Load(id)
	if !ok {
		return prod, utils.Error(utils.ProductNotFound, id)
	}

	prod, ok = result.(models.Product)
	if !ok {
		return prod, utils.Error(utils.ProductNotFound, id)
	}
	return prod, nil
}

func (p *ProductsDB) Upsert(product models.Product) {
	p.products.Store(product.ID, product)
}

func (p *ProductsDB) FindAll() []models.Product {
	result := make([]models.Product, 0)

	p.products.Range(func(key, value interface{}) bool {
		result = append(result, value.(models.Product))
		return true
	})

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}
