package service

import (
	"exam/exam/dao"
	"exam/exam/model"
)

func CreateUser(u model.User) (err error) {
	err = dao.CreateUser(u)
	return
}
