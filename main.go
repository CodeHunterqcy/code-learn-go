package main

import (
	"code-learn-go/hot100/mid"
	"fmt"
)

func main() {
	topKFrequent := mid.TopKFrequent([]int{1, 1, 1, 2, 2, 2, 3, 3, 4, 4, 4, 5}, 2)
	fmt.Print(topKFrequent)
}
