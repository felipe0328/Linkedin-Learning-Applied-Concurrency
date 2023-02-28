package controller

import (
	"appliedConcurrency/db"
	"appliedConcurrency/models"
	"appliedConcurrency/utils"
	"math"
	"sync"
)

type IController interface {
	CreateOrder(item models.Item) (*models.Order, error)
	GetAllProducts() []models.Product
	GetOrder(id string) (models.Order, error)
}

type controller struct {
	sync.Mutex
	productDB db.IProductsDB
	ordersDB  db.IOrderDB
}

func NewController(products db.IProductsDB, orders db.IOrderDB) IController {
	return &controller{
		productDB: products,
		ordersDB:  orders,
	}
}

func (c *controller) CreateOrder(item models.Item) (*models.Order, error) {
	err := c.validateItem(item)
	if err != nil {
		return nil, err
	}

	order := models.NewOrder(item)
	c.ordersDB.Upsert(order)

	go c.processNewOrder(&order)
	return &order, nil
}

func (c *controller) GetAllProducts() []models.Product {
	result := c.productDB.FindAll()
	return result
}

func (c *controller) GetOrder(id string) (models.Order, error) {
	return c.ordersDB.Find(id)
}

func (c *controller) validateItem(item models.Item) error {
	if item.Amount < 1 {
		return utils.Error(utils.OrderAmountWrong, item.Amount)
	}

	if exists := c.productDB.Exists(item.ProductID); !exists {
		return utils.Error(utils.ProductNotExists, item.ProductID)
	}

	return nil
}

func (c *controller) processNewOrder(order *models.Order) {
	c.Lock()
	defer c.Unlock()
	c.processOrder(order)
	c.ordersDB.Upsert(*order)
}

func (c *controller) processOrder(order *models.Order) {
	item := order.Item
	product, err := c.productDB.Find(item.ProductID)
	if err != nil {
		order.Status = models.OrderStatus_Rejected
		order.Error = err.Error()
		return
	}

	if product.Stock < item.Amount {
		order.Status = models.OrderStatus_Rejected
		order.Error = utils.Error(utils.ProductNotStock, item.ProductID, product.Stock, item.Amount).Error()
		return
	}

	product.Stock = product.Stock - item.Amount
	c.productDB.Upsert(product)

	total := math.Round(float64(order.Item.Amount)*product.Price*100) / 100
	order.Total = &total
	order.Complete()
}
