package controllers

import (
	"customer_service/configs"
	"customer_service/models"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateOrder(gContext *gin.Context) {
	var order models.Order
	err := gContext.ShouldBindJSON(&order)
	if err != nil {
		log.Printf("error in parsing order input:%s", err.Error())
		processError(gContext, configs.INVALID_INPUT, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	err = models.GetDb().CreateOrder(&order)
	if err != nil {
		if err == sql.ErrNoRows {
			processError(gContext, configs.INVALID_PRODUCT_ID, http.StatusBadRequest, configs.STATUS_ERROR)
		} else if strings.Contains(err.Error(), "expired") {
			processError(gContext, configs.DEAL_EXPIRED, http.StatusOK, configs.STATUS_ERROR)
		} else {
			log.Printf("error in creating order:%s", err.Error())
			processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		}
		return
	}
	response := make(Response)
	response[configs.KEY_STATUS] = configs.STATUS_SUCCESS
	response[configs.KEY_MSG] = gin.H{"order_details": order}
	gContext.JSON(http.StatusOK, response)
}
func OrderStatus(gContext *gin.Context) {
	orderIDStr := gContext.Param("order_id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		log.Printf("erorr in converting order-id to int:%s", err.Error())
		processError(gContext, configs.INVALID_ORDER_ID, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}

	orderStatus, err := models.GetDb().GetOrderStatus(int64(orderID))
	if err != nil {

		if err == sql.ErrNoRows {
			processError(gContext, configs.INVALID_ORDER_ID, http.StatusBadRequest, configs.STATUS_ERROR)
			return
		}
		log.Printf("error in getting order status:%s", err.Error())
		processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		return
	}
	response := make(Response)
	response[configs.KEY_STATUS] = configs.STATUS_SUCCESS
	response[configs.KEY_MSG] = gin.H{"order_status": orderStatus}
	gContext.JSON(http.StatusOK, response)

}
func Products(gContext *gin.Context) {

	products, err := models.GetDb().GetUnExpiredProductDeal()
	if err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				processError(gContext, configs.NO_DEAL_PRESENT, http.StatusOK, configs.STATUS_SUCCESS)
				return
			}
			log.Printf("error in fetcing un-expired deals:%s", err.Error())
			processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
			return
		}
	}
	response := make(Response)
	response[configs.KEY_STATUS] = configs.STATUS_SUCCESS
	response[configs.KEY_MSG] = gin.H{"products": products}
	gContext.JSON(http.StatusOK, response)
}
