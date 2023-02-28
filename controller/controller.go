package controller

import (
	"appliedConcurrency/db"
	"appliedConcurrency/models"
	"appliedConcurrency/stats"
	"appliedConcurrency/utils"
	"context"
	"math"
	"sync"
)

type IController interface {
	CreateOrder(item models.Item) (*models.Order, error)
	GetAllProducts() []models.Product
	GetOrder(id string) (models.Order, error)
	CloseOrders()
	GetOrderStats(context.Context) models.Statistics
}

type controller struct {
	sync.Mutex
	productDB db.IProductsDB
	ordersDB  db.IOrderDB
	stats     stats.IStatsService
	incoming  chan models.Order
	processed chan models.Order
	done      chan struct{}
}

func NewController(products db.IProductsDB, orders db.IOrderDB) IController {
	processedChan := make(chan models.Order)
	doneChan := make(chan struct{})

	stats := stats.NewStats(processedChan, doneChan)

	c := &controller{
		productDB: products,
		ordersDB:  orders,
		stats:     stats,
		incoming:  make(chan models.Order),
		processed: processedChan,
		done:      doneChan,
	}

	go c.processOrders()

	return c
}

func (c *controller) CreateOrder(item models.Item) (*models.Order, error) {
	err := c.validateItem(item)
	if err != nil {
		return nil, err
	}

	order := models.NewOrder(item)

	select {
	case c.incoming <- order:
		c.ordersDB.Upsert(order)
		return &order, nil
	case <-c.done:
		return nil, utils.Error(utils.OrdersClosed)
	}
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

func (c *controller) processOrders() {
	for {
		select {
		case order := <-c.incoming:
			c.processOrder(&order)
			c.ordersDB.Upsert(order)
			c.processed <- order
		case <-c.done:
			return
		}
	}
	// for order := range c.incoming {
	// 	c.processOrder(&order)
	// 	c.ordersDB.Upsert(order)
	// }
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

func (c *controller) CloseOrders() {
	close(c.done)
}

func (c *controller) GetOrderStats(ctx context.Context) (models.Statistics, error) {
	select {
	case s := <-c.stats.GetStats():
		return s, nil
	case <-ctx.Done():
		return models.Statistics{}, ctx.Err()
	}
}
