package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	Username string
	Nickname string
	Sex      uint8
	Birthday time.Time
}

// 结构体字段名字不是大写无法获取 反射
func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
