package waitgrouptest

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func process(i int) {
	defer wg.Done()
	fmt.Println("yes", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go process(i)
	}
	runtime.Gosched()

	wg.Wait()
	fmt.Print("main end")

}
