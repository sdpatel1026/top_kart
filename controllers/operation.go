package controllers

import (
	"Topkart/configs"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateProduct(gContext *gin.Context) {
	var input interface{}
	err := gContext.ShouldBindJSON(&input)
	if err != nil {
		fmt.Printf("error in parsing product input: %v\n", err.Error())
		processError(gContext, configs.INVALID_INPUT, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	url := fmt.Sprintf("%s/product", strings.TrimRight(configs.GetEnv("OPERATION_BASE_URL"), "/"))
	processRequest(gContext, input, url, http.MethodPost)

}
func UpdateProduct(gContext *gin.Context) {
	var input interface{}
	err := gContext.ShouldBindJSON(&input)
	if err != nil {
		fmt.Printf("error in parsing product input: %v\n", err.Error())
		processError(gContext, configs.INVALID_INPUT, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	fmt.Printf("input: %v\n", input)
	url := fmt.Sprintf("%s/product", strings.TrimRight(configs.GetEnv("OPERATION_BASE_URL"), "/"))
	processRequest(gContext, input, url, http.MethodPut)
}
func ApproveOrder(gContext *gin.Context) {
	orderID := gContext.Param("order_id")
	url := fmt.Sprintf("%s/approve/%s", OPERATION_BASE_URL, orderID)
	processRequest(gContext, nil, url, http.MethodPatch)
}
