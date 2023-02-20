package controllers

import (
	"customer_service/configs"

	"github.com/gin-gonic/gin"
)

type Response map[string]interface{}

func processError(gContext *gin.Context, errMsg string, httpStausCode, statusCode int) {
	response := make(Response)
	response[configs.KEY_STATUS] = statusCode
	response[configs.KEY_MSG] = errMsg
	gContext.JSON(httpStausCode, response)
}
