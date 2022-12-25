package util

import "github.com/gin-gonic/gin"

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var ok = respTemplate{
	Status: 200,
	Info:   "ok",
}
var internalError = respTemplate{
	Status: 500,
	Info:   "ok",
}
var ParamError = respTemplate{
	Status: 300,
	Info:   "param error",
}

func RespParamError(c *gin.Context) {
	c.JSON(300, ParamError)
}
func RespOK(c *gin.Context) {
	c.JSON(200, ok)
}
func RespInternalError(c *gin.Context) {
	c.JSON(500, internalError)
}
func AnyInfo(c *gin.Context, status int, info string) {
	c.JSON(status, gin.H{
		"info": info,
	})
}
