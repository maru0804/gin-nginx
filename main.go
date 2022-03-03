package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maru0804/gin-nginx/hello"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/user", hello.Hello())
	}

	router.Run(":9000")
}
