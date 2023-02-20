package controllers

import (
	"operation_service/configs"
	"time"

	"github.com/gin-gonic/gin"
)

type Response map[string]interface{}

var maxExpiryTime time.Time

func init() {
	maxExpiryTime, _ = time.Parse("15:04:05", "12:00:00")
}
func processError(gContext *gin.Context, errMsg string, httpStausCode, statusCode int) {
	response := make(Response)
	response[configs.KEY_STATUS] = statusCode
	response[configs.KEY_MSG] = errMsg
	gContext.JSON(httpStausCode, response)
}
