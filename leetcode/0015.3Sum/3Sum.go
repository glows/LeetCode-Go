package leetcode

import (
	"sort"
)

// 解法一 最优解， 双指针 + 排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, index-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index + 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[index] + nums[start] + nums[end]
			if addNum == 0 {
				result = append(result, []int{nums[index], nums[start], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}

		}
	}
	return result
}
