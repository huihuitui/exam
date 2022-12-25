package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/login", Login)
	house := r.Group("/house")
	{
		house.POST("/in", In)
		house.POST("/out", Out)
	}
}
