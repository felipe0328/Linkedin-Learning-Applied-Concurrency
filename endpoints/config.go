package endpoints

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const serverPort = ":3000"

func PrepareAndStartServer() error {
	router := gin.Default()
	handler, err := newHandler()

	if err != nil {
		return err
	}

	configureEndpoints(router, handler)

	fmt.Printf("Starting server on localhost%s", serverPort)
	log.Fatal(router.Run(serverPort))

	return nil
}

func configureEndpoints(router *gin.Engine, handler iHandler) error {
	router.GET("/", handler.GetIndex)
	router.GET("/products", handler.GetProducts)
	router.GET("/orders/:orderId", handler.GetOrderByID)
	router.POST("/orders", handler.CreateNewOrder)

	return nil
}
