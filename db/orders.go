package db

import (
	"appliedConcurrency/models"
	"appliedConcurrency/utils"
	"sync"
)

type IOrderDB interface {
	Find(id string) (models.Order, error)
	Upsert(order models.Order)
}

type OrderDB struct {
	placedOrders sync.Map
}

func NewOrders() IOrderDB {
	return &OrderDB{}
}

func (o *OrderDB) Find(id string) (models.Order, error) {
	var order models.Order
	result, ok := o.placedOrders.Load(id)
	if !ok {
		return order, utils.Error(utils.OrderNotFound, id)
	}

	order, ok = result.(models.Order)
	if !ok {
		return order, utils.Error(utils.OrderNotFound, id)
	}

	return order, nil
}

func (o *OrderDB) Upsert(order models.Order) {
	o.placedOrders.Store(order.ID, order)
}

func (o *OrderDB) Delete(order models.Order) {
	o.placedOrders.Delete(order.ID)
}
