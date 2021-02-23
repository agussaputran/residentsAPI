package controllers

import "github.com/gin-gonic/gin"

// ResultAPIResponse func
func ResultAPIResponse(str interface{}, length int) gin.H {
	return gin.H{
		"code":    200,
		"status":  "Ok",
		"message": "success get data",
		"data":    str,
		"count":   length,
	}
}

// ResultAPINilResponse func
func ResultAPINilResponse(str interface{}, length int) gin.H {
	return gin.H{
		"code":    200,
		"status":  "Ok",
		"message": "success get data",
		"data":    nil,
		"count":   length,
	}
}
