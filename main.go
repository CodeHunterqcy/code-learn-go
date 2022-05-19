package main

import (
	"code-learn-go/hot100/mid"
	"fmt"
)

func main() {
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	order := mid.SpiralOrder(arr)
	fmt.Print(order)

}
