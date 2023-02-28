package endpoints

import (
	"appliedConcurrency/controller"
	"appliedConcurrency/models"
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type iHandler interface {
	GetIndex(*gin.Context)
	GetProducts(*gin.Context)
	GetOrderByID(*gin.Context)
	CreateNewOrder(*gin.Context)
	Close(*gin.Context)
	GetStats(*gin.Context)
	RequestReversal(*gin.Context)
}

type handler struct {
	c controller.IController
	sync.Once
}

func newHandler(cont controller.IController) (iHandler, error) {
	h := handler{
		c: cont,
	}
	return &h, nil
}

func (h *handler) GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the orders App!")
}

func (h *handler) GetProducts(c *gin.Context) {
	products := h.c.GetAllProducts()
	c.JSON(http.StatusOK, products)
}

func (h *handler) GetOrderByID(c *gin.Context) {
	id := c.Param("orderID")
	order, err := h.c.GetOrder(id)

	if err != nil {
		c.String(http.StatusNotAcceptable, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *handler) CreateNewOrder(c *gin.Context) {
	var item models.Item

	err := c.BindJSON(&item)
	if err != nil {
		c.String(http.StatusNotAcceptable, err.Error())
		return
	}

	newModel, err := h.c.CreateOrder(item)

	if err != nil {
		c.String(http.StatusNotAcceptable, err.Error())
		return
	}

	c.JSON(http.StatusOK, newModel)
}

func (h *handler) Close(c *gin.Context) {
	h.Once.Do(func() {
		h.c.CloseOrders()
	})
	c.String(http.StatusAccepted, "The orders has been closed.")
}

func (h *handler) GetStats(c *gin.Context) {
	reqCtx := c.Request.Context()
	ctx, cancel := context.WithTimeout(reqCtx, 100*time.Millisecond)
	defer cancel()
	stats, err := h.c.GetOrderStats(ctx)
	if err != nil {
		c.String(http.StatusNotAcceptable, err.Error())
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *handler) RequestReversal(c *gin.Context) {
	orderID := c.Param("orderID")
	order, err := h.c.RequestReversal(orderID)
	if err != nil {
		c.String(http.StatusNotAcceptable, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, order)
}
