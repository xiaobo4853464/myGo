package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	defer func() {
		engine.Run(":8093")
	}()

	rg := engine.Group("/user")
	rg.POST("/register")
	rg.POST("/login", registerHandle)
	rg.GET("/userinfo", func(context *gin.Context) {
		path := context.FullPath()
		context.String(http.StatusOK, path)
	})
	rg.DELETE("/:id", func(context *gin.Context) {
		path := context.FullPath()
		uid := context.Param("id")
		fmt.Println("要删除的uid:", uid)
		context.String(http.StatusOK, path)
	})
}

func registerHandle(context *gin.Context) {
	path := context.FullPath()
	context.String(http.StatusOK, path)
}
