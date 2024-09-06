package router

import (
	"goAssignmentProject/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("/orders", controller.CreateOrder)
		orderGroup.GET("/orders", controller.GetOrders)
		orderGroup.GET("/orders/:id", controller.GetOrder)
		orderGroup.PUT("/orders/:id", controller.UpdateOrder)
		orderGroup.DELETE("/orders/:id", controller.DeleteOrder)
	}

	return r
}
