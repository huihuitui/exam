package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/login", Login)
	r.POST("/house", Add)
}
