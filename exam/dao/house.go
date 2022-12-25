package dao

import "exam/exam/model"

func Select(id int, name string) (g model.Good, err error) {
	Row := DB.QueryRow("select * from house where good_name=? && huose_id=?", id, name)
	if err = Row.Err(); Row.Err() != nil {
		return
	}
	err = Row.Scan(&g.HouseID, &g.GoodName, &g.GoodNum)
	return
}
func InGood(g model.Good) (err error) {
	_, err = DB.Exec("insert into house (good_name ,good_num,house_id)values (?.?,?)", g.GoodName, g.GoodNum, g.HouseID)
	return
}
func UpdateInGood(g model.Good, num int) (err error) {
	_, err = DB.Exec("insert into house (good_name ,good_num,house_id)values (?.?,?)", g.GoodName, num+g.GoodNum, g.HouseID)
	return
}
func OutGood(g model.Good) (err error) {
	_, err = DB.Exec("insert into house (good_name ,good_num,house_id)values (?.?,?)", g.GoodName, g.GoodNum, g.HouseID)
	return
}
func UpdateOutGood(g model.Good, num int) (err error) {
	_, err = DB.Exec("insert into house (good_name ,good_num,house_id)values (?.?,?)", g.GoodName, num-g.GoodNum, g.HouseID)
	return
}
