package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch chan int

func main() {
	ch = make(chan int, 1)
	wg.Add(3)
	go Work("goroutine1")
	go Work("goroutine2")
	go Work("goroutine3")
	wg.Wait()
	fmt.Println("successful")
}

func Work(workName string) {
	ch <- 1
	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	wg.Done()
	<-ch
}
