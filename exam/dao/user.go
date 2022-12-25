package dao

import "exam/exam/model"

func CreateUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user (user_name ,password) values (?,?)", u.UserName, u.Password)
	return err
}
