package syntest

import (
	"fmt"
	"sync"
	"time"
)

/*
假设我们有一个固定长度为2的队列，并且我们要将10个元素放入队列中。 我们希望一有空间就能放入，所以在队列中有空间时需要立刻通知
*/

func learnCond() {
	c := sync.NewCond(&sync.Mutex{})    //1
	queue := make([]interface{}, 0, 10) //2

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        //8
		queue = queue[1:] //9
		fmt.Println("Removed from queue")
		c.L.Unlock() //10
		c.Signal()   //11
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            //3
		for len(queue) == 2 { //4
			c.Wait() //5
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) //6
		c.L.Unlock()                        //7
	}
}
