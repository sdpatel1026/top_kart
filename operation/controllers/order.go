package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"operation_service/configs"
	"operation_service/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApproveOrder(gContext *gin.Context) {
	orderIDStr := gContext.Param("order_id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		log.Printf("erorr in converting order-id to int:%s", err.Error())
		processError(gContext, configs.INVALID_ORDER_ID, http.StatusBadRequest, configs.STATUS_ERROR)
		return
	}
	err = models.GetDb().ApproveOrder(orderID)
	if err != nil {
		errMsg := err.Error()
		if err == sql.ErrNoRows {
			processError(gContext, configs.INVALID_ORDER_ID, http.StatusBadRequest, configs.STATUS_ERROR)
		} else if strings.Contains(errMsg, models.STATUS_REJECTES) {
			processError(gContext, configs.OREDER_REJECTED, http.StatusBadRequest, configs.STATUS_ERROR)
		} else if strings.Contains(errMsg, models.ITEM_SORTAGE) {
			processError(gContext, configs.NO_SUFFICIENT_ITEM, http.StatusOK, configs.STATUS_ERROR)
		} else {
			log.Printf("error in approving order:%s", errMsg)
			processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		}
		return
	}
	response := make(Response)
	response[configs.KEY_STATUS] = configs.STATUS_SUCCESS
	response[configs.KEY_MSG] = gin.H{"result": configs.ORDER_SUCCESSFULLY_APPROVED}
	gContext.JSON(http.StatusOK, response)
}
