package controllers

import (
	"Topkart/configs"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateOrder(gContext *gin.Context) {
	var input interface{}
	err := gContext.ShouldBindJSON(&input)
	if err != nil {
		fmt.Printf("error in parsing order input: %v\n", err.Error())
		processError(gContext, configs.INVALID_INPUT, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	region := strings.TrimSpace(gContext.GetHeader("region"))
	if region == "" {
		processError(gContext, configs.REGION_MISSING, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	} else if !configs.REGIONS[region] {
		processError(gContext, configs.INVALID_REGION, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	var url string
	if region == "AR" {
		url = fmt.Sprintf("%s/order", CUSTOMER_BASE_URL_AR)
	} else if region == "AFR" {
		url = fmt.Sprintf("%s/order", CUSTOMER_BASE_URL_AFR)
	} else if region == "AMR" {
		url = fmt.Sprintf("%s/order", CUSTOMER_BASE_URL_AMR)
	} else {
		url = fmt.Sprintf("%s/order", CUSTOMER_BASE_URL_EUR)
	}
	processRequest(gContext, input, url, http.MethodPost)
}

func OrderStaus(gContext *gin.Context) {

	orderID := gContext.Param("order_id")
	if orderID == "" {
		processError(gContext, configs.ORDER_ID_MISSING, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	region := strings.TrimSpace(gContext.GetHeader("region"))
	if region == "" {
		processError(gContext, configs.REGION_MISSING, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	} else if !configs.REGIONS[region] {
		processError(gContext, configs.INVALID_REGION, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	var url string
	if region == "AR" {
		url = fmt.Sprintf("%s/order/status/%s", CUSTOMER_BASE_URL_AR, orderID)
	} else if region == "AFR" {
		url = fmt.Sprintf("%s/order/stauts/%s", CUSTOMER_BASE_URL_AFR, orderID)
	} else if region == "AMR" {
		url = fmt.Sprintf("%s/order/status/%s", CUSTOMER_BASE_URL_AMR, orderID)
	} else {
		url = fmt.Sprintf("%s/order/status/%s", CUSTOMER_BASE_URL_EUR, orderID)
	}
	processRequest(gContext, nil, url, http.MethodGet)
}

func GetProducts(gContext *gin.Context) {
	region := strings.TrimSpace(gContext.GetHeader("region"))
	if region == "" {
		processError(gContext, configs.REGION_MISSING, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	} else if !configs.REGIONS[region] {
		processError(gContext, configs.INVALID_REGION, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	var url string
	if region == "AR" {
		url = fmt.Sprintf("%s/products", CUSTOMER_BASE_URL_AR)
	} else if region == "AFR" {
		url = fmt.Sprintf("%s/products", CUSTOMER_BASE_URL_AFR)
	} else if region == "AMR" {
		url = fmt.Sprintf("%s/products", CUSTOMER_BASE_URL_AMR)
	} else {
		url = fmt.Sprintf("%s/products", CUSTOMER_BASE_URL_EUR)
	}
	fmt.Printf("url: %v\n", url)
	processRequest(gContext, nil, url, http.MethodGet)
}
