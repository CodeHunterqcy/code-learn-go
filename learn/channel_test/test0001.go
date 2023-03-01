package channel_test

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	用 goroutine 和 channel 协作完成
	1、开启一个writeData协程，向管道intChan中写入50个整数
	2、开启一个readData协程，从管道intChan中读取writeData写入的数据

*/

var intChan chan int

func test0001() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	if <-exitChan {
		return
	}

}
func writeData(intChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		intChan <- rand.Intn(4)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		val, ok := <-intChan
		if !ok {
			break
		}
		fmt.Print(val)
	}
	exitChan <- true
	close(intChan)
}
