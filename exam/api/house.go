package api

import (
	"database/sql"
	"exam/exam/model"
	"exam/exam/service"
	"exam/exam/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Add(c *gin.Context) {

}
func In(c *gin.Context) {
	goodName := c.PostForm("good_name")
	goodNum := c.PostForm("good-num")
	houseID := c.PostForm("house_id")
	id, err := strconv.Atoi(houseID)
	if err != nil {
		fmt.Println(err)
		return
	}
	num, err := strconv.Atoi(goodNum)
	if err != nil {
		fmt.Println(err)
		return
	}
	good, err := service.Select(id, goodName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("errererer:%v", err)
		util.RespInternalError(c)
		return
	}
	if err == sql.ErrNoRows {
		err = service.InGood(model.Good{
			GoodName: goodName,
			GoodNum:  num,
			HouseID:  id,
		})
		if err != nil {
			util.RespInternalError(c)
			return
		}
		util.RespOK(c)
	}
	err = service.UpdateInGood(model.Good{
		GoodName: goodName,
		GoodNum:  num,
		HouseID:  id,
	}, good.GoodNum)
	if err != nil {
		util.RespInternalError(c)
		return
	}
	util.RespOK(c)
}
func Out(c *gin.Context) {
	goodName := c.PostForm("good_name")
	goodNum := c.PostForm("good-num")
	houseID := c.PostForm("house_id")
	id, err := strconv.Atoi(houseID)
	if err != nil {
		fmt.Println(err)
		return
	}
	num, err := strconv.Atoi(goodNum)
	if err != nil {
		fmt.Println(err)
		return
	}
	good, err := service.Select(id, goodName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("errererer:%v", err)
		util.RespInternalError(c)
		return
	}
	if err == sql.ErrNoRows {
		err = service.OutGood(model.Good{
			GoodName: goodName,
			GoodNum:  num,
			HouseID:  id,
		})
		if err != nil {
			util.RespInternalError(c)
			return
		}
		util.RespOK(c)
	}
	err = service.UpdateOutGood(model.Good{
		GoodName: goodName,
		GoodNum:  num,
		HouseID:  id,
	}, good.GoodNum)
	if err != nil {
		util.RespInternalError(c)
		return
	}
	util.RespOK(c)
}
