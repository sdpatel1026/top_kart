package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("hostAddr and appPort required")
	}
	hostAddr := os.Args[1]
	appPort := os.Args[2]
	fmt.Printf("hostAddr: %v\n", hostAddr)
	fmt.Printf("appPort: %v\n", appPort)
	setupRoutes()
	err := router.Run(fmt.Sprintf("%s:%s", hostAddr, appPort))
	if err != nil {
		log.Fatalf("error in starting service:%s", err.Error())
	}
}
