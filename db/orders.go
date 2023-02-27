package db

import (
	"appliedConcurrency/models"
	"appliedConcurrency/utils"
)

type IOrderDB interface {
	Find(id string) (models.Order, error)
	Upsert(order models.Order)
}

type OrderDB struct {
	placedOrders map[string]models.Order
}

func NewOrders() IOrderDB {
	return &OrderDB{
		placedOrders: make(map[string]models.Order),
	}
}

func (o *OrderDB) Find(id string) (models.Order, error) {
	var order models.Order
	order, ok := o.placedOrders[id]
	if !ok {
		return order, utils.Error(utils.OrderNotFound)
	}

	return order, nil
}

func (o *OrderDB) Upsert(order models.Order) {
	o.placedOrders[order.ID] = order
}
