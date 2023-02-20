package main

import (
	"customer_service/controllers"

	"github.com/gin-gonic/gin"
)

func setupRoutes() {
	router = gin.New()
	router.GET("/order/status/:order_id", controllers.OrderStatus)
	router.GET("/products", controllers.Products)
	router.POST("/order", controllers.CreateOrder)
}
