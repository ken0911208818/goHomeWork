package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ken0911208818/goHomeWork/handler"
)

func main() {
	r := setupRouter()
	r.Run("")
}

func setupRouter() *gin.Engine{
	r := gin.Default()
	role := r.Group("/role")
	{
		role.GET("/", handler.Index)
		role.POST("/", handler.Create)
		role.GET("/:id", handler.GetOne)
		role.PUT("/:id", handler.Update)
		role.DELETE("/:id", handler.Delete)
	}
	return r
}