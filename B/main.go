package main

import "fmt"

// if 条件语句中return 最后面的打印3得语句没执行到 所以也不会有打印3的操作
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()

	if a {
		fmt.Println("2")

	}

	defer func() {
		fmt.Println("3")
	}()
}
