package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 計算1-200的各個數的階乘，並且把各個數的階乘放到map中
// 最後顯示出來，使用goroutine完成

// 思路
// 1. 編寫一個函數，來計算各個數的階乘，並放到map中
// 2. 啟動多個協程，並將統計的結果放到map中
// 3. 使用一個全局的map

var (
	myMap = make(map[int]int, 10)
	// 聲明一個全局的互斥鎖
	// lock是一個全局的互斥鎖
	// sync是包: synchornized 同步
	lock sync.Mutex
)

// test函數就是計算 n! ，並將結果放到myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 這裡我們將res放入到myMap
	// 加鎖
	lock.Lock()
	myMap[n] = res
	// 解鎖
	lock.Unlock()
}

func main() {
	// 查看系統的CPU核心數
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum = ", cpuNum)

	// 可以自己設置使用多個cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")

	// 我們這裡開啟多個協程，來完成任務
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	// 休眠5秒鐘，讓主程序不要那麼快結束
	//time.Sleep(time.Second * 5)

	// 輸出結果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}
