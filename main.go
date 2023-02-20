package main

import (
	"Topkart/configs"
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	setupRoutes()
	appPort := configs.GetEnv("APP_PORT")
	hostAddr := configs.GetEnv("HOST_ADDR")
	router.Run(fmt.Sprintf("%s:%s", hostAddr, appPort))
}
