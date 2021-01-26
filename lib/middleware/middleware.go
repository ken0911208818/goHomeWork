package middleware

import "github.com/gin-gonic/gin"

func SendResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func SendErrorResponse(c *gin.Context, StatusCode int, err error) {
	c.JSON(StatusCode, err.Error())
}
