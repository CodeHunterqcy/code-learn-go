package hard

/*
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。
https://leetcode.cn/problems/median-of-two-sorted-arrays/?favorite=2cktkvj
*/

// findMedianSortedArrays 合并数组版
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res []int
	len1, len2 := len(nums1), len(nums2)
	l1, l2 := 0, 0
	for l1 < len1 && l2 < len2 {
		if nums1[l1] < nums2[l2] {
			res = append(res, nums1[l1])
			l1++
		} else {
			res = append(res, nums2[l2])
			l2++
		}
	}
	res = append(res, nums1[l1:]...)
	res = append(res, nums2[l2:]...)
	length := len1 + len2
	if length%2 == 1 {
		return float64(res[length/2])
	}

	return float64(res[length/2]+res[length/2-1]) / 2.0
}

// findMedianSortedArrays2 双指针寻找第K小的数字
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	l1, l2, res1, res2 := 0, 0, 0, 0
	length := len1 + len2

	for i := 0; i <= length/2; i++ {
		res1 = res2
		if l1 < len1 && (l2 >= len2 || nums1[l1] < nums2[l2]) {
			res2 = nums1[l1]
			l1++
		} else {
			res2 = nums2[l2]
			l2++
		}
	}
	if length%2 == 1 {
		return float64(res2)
	}

	return float64(res1+res2) / 2.0
}

// findMedianSortedArrays1 答案版-二分查找寻找第K小数字
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return float64(findK(length/2+1, nums1, nums2))
	}
	k1, k2 := length/2, length/2+1
	return float64(findK(k1, nums1, nums2)+findK(k2, nums1, nums2)) / 2.0
}

// findK 两个有序数组二分寻找第K小数字
func findK(k int, nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	index1, index2 := 0, 0
	for {
		if m == index1 {
			return nums2[index2+k-1]
		}
		if n == index2 {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return getMin(nums1[index1], nums2[index2])
		}

		half := k / 2
		l1 := getMin(index1+half, m) - 1
		l2 := getMin(index2+half, m) - 1
		pivot1, pivot2 := nums1[l1], nums2[l2]
		if pivot1 <= pivot2 {
			k -= l1 - index1 + 1
			index1 = l1 + 1
		} else {
			k -= l2 - index2 + 1
			index2 = l2 + 1
		}
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
