package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	defer func() {
		if err := engine.Run(":8092"); err != nil {
			log.Fatal(err.Error())
		}
	}()
	//engine.GET("/hello", func(context *gin.Context) {
	//	fmt.Println("The full path is :", context.FullPath())
	//	context.Writer.Write([]byte("111\n"))
	//})

	//使用handle 处理get
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		name := context.DefaultQuery("name", "nil")
		fmt.Println(name)
		//context.Writer.Write([]byte("HELLO " + name))
		context.String(http.StatusOK, "hello: "+name)

	})

	//使用handle 处理post
	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		usr := context.PostForm("username")
		//context.Param("username")
		//context.Param("")
		pwd := context.PostForm("password")
		fmt.Println(usr, pwd)
		context.Writer.Write([]byte(usr + " login"))

	})

	//使用内置delete 处理
	engine.DELETE("user/:id", func(context *gin.Context) {
		userId := context.Param("id")
		fmt.Println("userId:", userId)
		context.Writer.Write([]byte("Delete user id:" + userId))
	})

	// 数据绑定，get
	engine.GET("/userinfo", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var u userRegister
		if err := context.ShouldBindQuery(&u); err != nil {
			log.Fatal(err.Error())
			return
		}

		words := fmt.Sprintf("username:%s,gender:%s,address:%s\n", u.Name, u.Gender, u.Addr)
		fmt.Println(words)
		context.Writer.Write([]byte(words))

	})

	// 数据绑定by json，post
	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var u userRegister
		if err := context.BindJSON(&u); err != nil {
			log.Fatal(err.Error())
			return
		}

		words := fmt.Sprintf("username:%s,gender:%s,address:%s\n", u.Name, u.Gender, u.Addr)
		fmt.Println(words)
		context.Writer.Write([]byte(words))

	})

	//json return
	engine.GET("/json", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		context.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "ok",
			"detail":  context.FullPath(),
		})
	})

	//json struct return
	engine.POST("/jsons", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var user userRegister
		if err := context.BindJSON(&user); err != nil {
			log.Fatal(err.Error())
			return
		}
		context.JSON(http.StatusOK, user)
	})

}

type userRegister struct {
	Name   string `form:"name"`
	Gender string `form:"gender"`
	Addr   string `form:"addr"`
}
