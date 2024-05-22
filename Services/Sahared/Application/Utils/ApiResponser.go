package SharedUtils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data any, statusCode int) {
	c.Header("Content-Type", "application/json")
	var respose = map[string]interface{}{
		"message":     "Success",
		"status_code": statusCode,
		"data":        data,
	}
	c.JSON(statusCode, respose)
}

func Error(c *gin.Context, data any, statusCode int) {
	c.Header("Content-Type", "application/json")
	respose := make(map[string]interface{})
	respose["message"] = "Error"
	respose["status_code"] = statusCode
	respose["data"] = data
	jsonResponse, _ := json.Marshal(respose)
	c.JSON(statusCode, jsonResponse)
}
