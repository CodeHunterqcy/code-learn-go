package channel_test

import "fmt"

func test() {
	unBufferChannel := make(chan int)
	buffChannel := make(chan int, 10)

	unBufferChannel <- 10
	buffChannel <- 10

	data := <-unBufferChannel

	fmt.Print(data)
}
