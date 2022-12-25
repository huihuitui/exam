package service

import (
	"exam/exam/dao"
	"exam/exam/model"
)

func Select(id int, name string) (g model.Good, err error) {
	g, err = dao.Select(id, name)
	return
}
func InGood(g model.Good) (err error) {
	err = dao.InGood(g)
	return
}
func UpdateInGood(g model.Good, num int) (err error) {
	err = dao.UpdateInGood(g, num)
	return
}
func OutGood(g model.Good) (err error) {
	err = dao.InGood(g)
	return
}
func UpdateOutGood(g model.Good, num int) (err error) {
	err = dao.UpdateInGood(g, num)
	return
}
