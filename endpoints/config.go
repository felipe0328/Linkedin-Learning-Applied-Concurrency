package endpoints

import (
	"appliedConcurrency/controller"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const serverPort = ":3000"

func PrepareAndStartServer(c controller.IController) error {
	router := gin.Default()
	handler, err := newHandler(c)

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
	router.GET("/orders/:orderID", handler.GetOrderByID)
	router.POST("/orders", handler.CreateNewOrder)
	router.POST("/close", handler.Close)
	router.GET("/stats", handler.GetStats)

	return nil
}
