package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maru0804/gin-nginx/hello"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// nil を返すためエラーになる
		v1.GET("/user", hello.Hello())
		// うまくいく
		v1.GET("/hello", func(ctx *gin.Context) {
			ctx.String(200, "Hello World")
		})
	}

	router.Run(":9000")
}
