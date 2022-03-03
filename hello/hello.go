package hello

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Hello() gin.HandlerFunc {
	log.Println("hello")
	return nil
}
