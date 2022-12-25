package api

import (
	"exam/exam/model"
	"exam/exam/service"
	"exam/exam/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	userName := c.PostForm("username")
	Password := c.PostForm("password")
	if userName == "" || Password == "" {
		util.RespParamError(c)
		return
	}
	err := service.CreateUser(model.User{
		UserName: userName,
		Password: Password,
	})
	if err != nil {
		util.RespInternalError(c)
		return
	}
	c.SetCookie("user", userName, 3600, "/", "localhost", false, true)
	util.RespOK(c)
}
