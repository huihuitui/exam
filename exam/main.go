package main

import (
	"exam/exam/api"
	"exam/exam/dao"
)

func main() {
	api.InitRouter()
	dao.InitDB()
}
