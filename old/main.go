package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用傳統的方法，測試運行的時間
	start := time.Now().Unix()

	for num := 1; num <= 100000; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			// 將結果輸出
			//fmt.Printf("質數 = %d\n", res)
		}
	}
	end := time.Now().Unix()
	fmt.Printf("使用傳統方法耗時 = %v 秒\n", end-start)

}
