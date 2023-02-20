package main

import (
	"Topkart/controllers"

	"github.com/gin-gonic/gin"
)

func setupRoutes() {
	router = gin.New()
	customer := router.Group("customer")
	customer.POST("/order", controllers.CreateOrder)
	customer.GET("/order/status/:order_id", controllers.OrderStaus)
	customer.GET("/products", controllers.GetProducts)
	operation := router.Group("operation")
	operation.POST("/product", controllers.CreateProduct)
	operation.PUT("/product", controllers.UpdateProduct)
	operation.PATCH("/order/approve/:order_id", controllers.ApproveOrder)
}
