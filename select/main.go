package main

import (
	"fmt"
)

func main() {

	// 使用Select可以解決從管道取數據的阻塞問題

	// 定義一個管道，放入10個int數據
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 定義一個管道，放入5個string數據
	stringChan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 傳統的方法在遍歷管道時，如果不關閉會阻塞而導致deadlock
	// 在實際開發中，有時候不好確定什麼時候該關閉管道
	// 可以使用select方式來解決
	for {
		select {
		case v := <-intChan:
			fmt.Printf("從intChan讀取的數據%d\n", v)
		case v := <-stringChan:
			fmt.Printf("從stringChan讀取的數據%s\n", v)
		default:
			fmt.Printf("都取不到取值了啦!!!這邊可以加入邏輯!!!")
			return
		}
	}
}
