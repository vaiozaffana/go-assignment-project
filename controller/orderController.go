package controller

import (
	"goAssignmentProject/database"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order database.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order.OrderedAt = time.Now()
	order.UpdatedAt = time.Now()

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to create order"})
		return
	}
}

func GetOrders(c *gin.Context) {
	var orders []database.Order

	if err := database.DB.Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to get orders"})
		return
	}

	c.JSON(200, orders)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")

	var order database.Order
	if err := database.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	c.JSON(200, order)
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var order database.Order
	if err := database.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order.UpdatedAt = time.Now()

	if err := database.DB.Save(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to update order"})
		return
	}

	c.JSON(200, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	var order database.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	if err := database.DB.Delete(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to delete order"})
		return
	}

	c.JSON(200, gin.H{"message": "order deleted"})
}
