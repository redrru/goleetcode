package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	prevNumMap := map[int]int{}

	for index := range nums {
		if prevIndex, ok := prevNumMap[target-nums[index]]; ok {
			return []int{prevIndex, index}
		}

		prevNumMap[nums[index]] = index
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	require.Equal(t, []int{0, 1}, twoSum(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	require.Equal(t, []int{1, 2}, twoSum(nums, target))

	nums = []int{3, 3}
	target = 6
	require.Equal(t, []int{0, 1}, twoSum(nums, target))
}
