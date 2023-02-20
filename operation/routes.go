package main

import (
	"operation_service/controllers"

	"github.com/gin-gonic/gin"
)

func setupRoutes() {
	router = gin.New()
	router.POST("/product", controllers.CreateProduct)
	router.PUT("/product", controllers.UpdateProduct)
	router.PATCH("/approve/:order_id", controllers.ApproveOrder)
}
