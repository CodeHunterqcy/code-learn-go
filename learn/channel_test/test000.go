package channel_test

import (
	"fmt"
)

func test() {
	unBufferChannel := make(chan int)
	buffChannel := make(chan int, 10)

	unBufferChannel <- 10
	buffChannel <- 10

	data := <-unBufferChannel

	fmt.Print(data)

	// 遍历
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
	}()

	// 1
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	// 2
	for i := range c {
		fmt.Println(i)
	}

}
