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
	fmt.Println(search([]int{3, 1}, 1))

}
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		fmt.Printf("left=[%+v],right=[%+v],mid=[%+v]\n", left, right, mid)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] { //说明在前一段上
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
func longestValidParentheses(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if (j-i+1)%2 == 1 {
				continue
			}
			if isValid(s[i : j+1]) {
				if j-i+1 > res {
					res = j - i + 1
				}
			}
		}
	}
	return res
}
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	mapKey := map[byte]byte{
		')': '(',
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && mapKey[s[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
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

func testLock(i int) int {
	var myLock sync.Mutex
	myLock.Lock()
	i++
	myLock.Unlock()
	return i
}
