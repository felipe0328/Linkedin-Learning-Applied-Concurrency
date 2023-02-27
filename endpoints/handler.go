package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type iHandler interface {
	GetIndex(*gin.Context)
	GetProducts(*gin.Context)
	GetOrderByID(*gin.Context)
	CreateNewOrder(*gin.Context)
}

type handler struct{}

func newHandler() (iHandler, error) {
	h := handler{}
	return &h, nil
}

func (h *handler) GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the orders App!")
}

func (h *handler) GetProducts(c *gin.Context) {
}

func (h *handler) GetOrderByID(c *gin.Context) {
}

func (h *handler) CreateNewOrder(c *gin.Context) {
}
