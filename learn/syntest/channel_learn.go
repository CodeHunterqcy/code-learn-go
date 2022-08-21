package syntest

import (
	"fmt"
	"time"
)

func learnChannel() {
	c := make(chan int, 3)

	for i := 0; i < 10; i++ {
		go func() {
			c <- i
		}()
	}
	for c != nil {
		fmt.Print(<-c)
	}

}

func testChannel() {
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()
	fmt.Println("111")
	select {
	case <-c:
		fmt.Println(time.Since(start))
	}
}
