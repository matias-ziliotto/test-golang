package web

import "github.com/gin-gonic/gin"

type ResponseSuccess struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func Success(c *gin.Context, code int, data interface{}) {
	successResponse := ResponseSuccess{code, data}
	Response(c, code, successResponse)
}

func Error(c *gin.Context, code int, err string) {
	errResponse := ResponseError{code, err}
	c.AbortWithStatusJSON(code, errResponse)
}
