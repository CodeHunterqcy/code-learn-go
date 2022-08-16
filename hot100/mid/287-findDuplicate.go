package mid

import (
	"fmt"
)

// 287. 寻找重复数
/*
给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-the-duplicate-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func FindDuplicate(nums []int) int {
	// 连一条i-nums[i]的边，看是否有环
	fast, slow := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
		fmt.Printf("slow=[%+v],fast=[%+v]\n", slow, fast)
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
		fmt.Printf("slow=[%+v],fast=[%+v]\n", slow, fast)
	}
	return slow
}
func testFindDuplicate(nums []int) int {
	fast, slow := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
