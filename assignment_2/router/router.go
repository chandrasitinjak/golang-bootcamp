package router

import (
	"assignment_2/controller"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders", controller.GetOrder)
	router.GET("/order/:orderID", controller.GetOrderById)
	router.PUT("/orders/:orderID", controller.UpdateOrder)
	router.DELETE("/orders/:orderID", controller.DeleteOrder)

	return router
}
