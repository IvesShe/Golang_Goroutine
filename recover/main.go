package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello , world")
	}
}

func test() {
	// 這裡可以使用defer + recover
	defer func() {
		// 捕獲test拋出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 發生錯誤", err)
		}
	}()
	var myMay map[int]string
	myMay[0] = "Golang" //error
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok = ", i)
		time.Sleep(time.Second)
	}
}
