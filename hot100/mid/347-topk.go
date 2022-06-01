package mid

import "fmt"

// TopKFrequent top
/*
 前k个高频元素
输入: nums = [1,1,1,2,2,2,3,4,4,4,5], k = 3
*/
func TopKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	max := 0
	for _, num := range nums {
		numMap[num]++
		if numMap[num] > max {
			max = numMap[num]
		}
	}
	// numMap 1:3 2:3 3:1 4:2
	fmt.Println(numMap)
	list := make([][]int, max)
	for val, count := range numMap {
		list[count-1] = append(list[count-1], val)
	}
	// list: [3], [4], [1,2]
	fmt.Println(list)
	var res []int
	for i := max - 1; i >= 0; i-- {
		res = append(res, list[i]...)
		k = k - len(list[i])
		fmt.Printf("res=[%+v],k=[%+v].\n", res, k)
		if k == 0 {
			break
		}
	}

	return res

}

// 小顶堆
type IHeap [][2]int

func (h IHeap) len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
