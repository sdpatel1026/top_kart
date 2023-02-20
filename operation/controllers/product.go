package controllers

import (
	"log"
	"net/http"
	"operation_service/configs"
	"operation_service/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProduct(gContext *gin.Context) {
	var product models.Product
	err := gContext.ShouldBindJSON(&product)
	if err != nil {
		log.Printf("error in parsing input:%s", err.Error())
		// processError(gContext, err.Error(), http.StatusBadRequest, configs.STATUS_ERROR)
		processError(gContext, configs.INVALID_INPUT, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	expiryTime, err := time.Parse("15:04:05", product.ExpiryTime)
	if err != nil {
		log.Printf("error in parsing expiry time:%s", err.Error())
		processError(gContext, configs.INVALID_EXPIRY_TIME, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	if expiryTime.After(maxExpiryTime) {
		processError(gContext, configs.EXPIRTY_TIME_EXCEEDED, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	productID, err := models.GetDb().CreateProduct(product)
	if err != nil {
		log.Printf("error  in creating product:%s", err.Error())
		processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		return
	}
	response := make(Response)
	response[configs.KEY_STATUS] = configs.STATUS_SUCCESS
	response[configs.KEY_MSG] = gin.H{"product_id": productID}
	gContext.JSON(http.StatusOK, response)
}

func UpdateProduct(gContext *gin.Context) {
	var product models.Product
	err := gContext.ShouldBindJSON(&product)
	if err != nil {
		processError(gContext, configs.INVALID_INPUT, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}

	expiryTime, err := time.Parse("15:04:05", product.ExpiryTime)
	if err != nil {
		log.Printf("error in parsing expiry time:%s", err.Error())
		processError(gContext, configs.INVALID_EXPIRY_TIME, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	if expiryTime.After(maxExpiryTime) {
		processError(gContext, configs.EXPIRTY_TIME_EXCEEDED, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	if product.AvailableUnits > product.TotalUnits {
		processError(gContext, configs.INVALID_AVAILABLE_UNITS, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	if product.ID == 0 {
		processError(gContext, configs.PRODUCT_ID_MISSING, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	err = models.GetDb().UpdateProduct(product)
	if err != nil {
		processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		return
	}
	response := make(Response)
	response[configs.KEY_STATUS] = configs.STATUS_SUCCESS
	response[configs.KEY_MSG] = gin.H{"result": configs.PRODUCT_SUCCESSFULLY_UPDATED}
	gContext.JSON(http.StatusOK, response)
}
