package timer_test

import (
	"fmt"
	"time"
)

func testTimer() {
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Print(t1)

	t2 := <-timer1.C
	fmt.Println(t2)
}
