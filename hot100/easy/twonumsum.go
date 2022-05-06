package easy

// TwoSum1 暴力版
func TwoSum1(nums []int, target int) []int {
	if nums == nil {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// TwoSum 答案版
func TwoSum(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	// 用map存放target-num和下标
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := set[nums[i]]; ok {
			return []int{val, i}
		}
		set[target-nums[i]] = i
	}

	return nil
}
