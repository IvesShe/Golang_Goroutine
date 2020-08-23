package main

import (
	"fmt"
	"time"
)

// 向intChan放入數字
func putNum(intChan chan int) {

	for i := 1; i <= 100000; i++ {
		intChan <- i
	}

	// 關閉intChan
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		num, ok := <-intChan
		if !ok { // 當intChan取不到值
			break
		}

		// 判斷是不是質數
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 { // 質數只能被自己及1整除
				flag = false
				break
			}
		}
		if flag { // 將質數存入primeChan管道
			primeChan <- num
		}
	}

	fmt.Println("有一個primeNum協程因為取不到數據而退出!!")

	// 這裡不能關閉primeChan，因為可能還有其它primeNum使用
	// 向exitChan寫入true
	exitChan <- true
}

func main() {

	// 放入數字的管道
	intChan := make(chan int, 1000)

	// 放入質數結果的管道
	primeChan := make(chan int, 2000)

	// 標識退出的管道
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()

	// 開啟一個協程，向intChan放入數字
	go putNum(intChan)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 這邊跑一個匿名的協程，從exitChan取值
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Printf("使用協程耗時 = %v 秒\n", end-start)
		// 當取完4個值時，代表全部的協程都運行完畢了，可以關閉primeChan通道
		close(primeChan)
	}()

	// 遍歷primeNum，把結果取出
	for {
		//res, ok := <-primeChan
		_, ok := <-primeChan
		if !ok {
			break
		}
		// 將結果輸出
		//fmt.Printf("質數 = %d\n", res)
	}
}
