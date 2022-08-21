package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//全局变量
var ticket = 10 // 100张票

var wg sync.WaitGroup
var matex sync.Mutex // 创建锁头

func main() {
	///*
	//   4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	//*/
	//wg.Add(4)
	//go saleTickets("售票口1") // g1,100
	//go saleTickets("售票口2") // g2,100
	//go saleTickets("售票口3") //g3,100
	//go saleTickets("售票口4") //g4,100
	//wg.Wait()              // main要等待。。。
	//
	////time.Sleep(5*time.Second)
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} { //1
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done: //2
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() { //3
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated //4
	fmt.Println("Done.")
}

func countSubstrings(s string) int {
	length := len(s)
	if length == 0 || length == 1 {
		return length
	}

	res := 0
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if isHuiWen(s[j:i]) {
				res++
			}
		}
	}
	return res
}
func isHuiWen(s string) bool {
	start, end := 0, len(s)-1
	if start == end {
		return true
	}
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}

	return true
}
func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	//for i:=1;i<=100;i++{
	//  fmt.Println(name,"售出：",i)
	//}
	for { //ticket=1
		matex.Lock()
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket) // 1 , 0, -1 , -2
			ticket--                         //0 , -1 ,-2 , -3
		} else {
			matex.Unlock() //解锁
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
		matex.Unlock() //解锁
	}
}
