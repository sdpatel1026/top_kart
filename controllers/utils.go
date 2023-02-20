package controllers

import (
	"Topkart/configs"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var OPERATION_BASE_URL string = strings.TrimRight(configs.GetEnv("OPERATION_BASE_URL"), "/")
var CUSTOMER_BASE_URL_EUR string = strings.TrimRight(configs.GetEnv("CUSTOMER_BASE_URL_EUR"), "/")
var CUSTOMER_BASE_URL_AR string = strings.TrimRight(configs.GetEnv("CUSTOMER_BASE_URL_AR"), "/")
var CUSTOMER_BASE_URL_AFR string = strings.TrimRight(configs.GetEnv("CUSTOMER_BASE_URL_AFR"), "/")
var CUSTOMER_BASE_URL_AMR string = strings.TrimRight(configs.GetEnv("CUSTOMER_BASE_URL_AMR"), "/")

type Response map[string]interface{}

func processError(gContext *gin.Context, errMsg string, httpStausCode, statusCode int) {
	response := make(Response)
	response[configs.KEY_STATUS] = statusCode
	response[configs.KEY_MSG] = errMsg
	gContext.JSON(httpStausCode, response)
}

func processRequest(gContext *gin.Context, payload interface{}, url, method string) {

	marshalledInput, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error in marshalling input:%s", err.Error())
		processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		return
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(marshalledInput))
	if err != nil {
		log.Printf("error in making request:%s", err.Error())
		processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error in making request:%s", err.Error())
		processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		return
	}
	defer res.Body.Close()
	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("error in reading response:%s", err.Error())
		processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		return
	}
	var response Response
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		log.Printf("error in unmarshalling response:%s", err.Error())
		processError(gContext, configs.TECHNICAL_ERROR, http.StatusInternalServerError, configs.STATUS_ERROR)
		return
	}
	gContext.JSON(res.StatusCode, response)
}
