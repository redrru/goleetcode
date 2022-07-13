package array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, num := range nums {
		if pos, ok := seen[target-num]; ok {
			return []int{pos, i}
		}
		seen[num] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	type (
		args struct {
			Nums   []int
			Target int
		}
		want struct {
			Nums []int
		}
	)

	tests := []struct {
		Args args
		Want want
	}{
		{
			Args: args{
				Nums:   []int{2, 7, 11, 15},
				Target: 9,
			},
			Want: want{
				Nums: []int{0, 1},
			},
		},
		{
			Args: args{
				Nums:   []int{3, 2, 4},
				Target: 6,
			},
			Want: want{
				Nums: []int{1, 2},
			},
		},
		{
			Args: args{
				Nums:   []int{3, 3},
				Target: 6,
			},
			Want: want{
				Nums: []int{0, 1},
			},
		},
	}

	for _, tc := range tests {
		require.Equal(t, tc.Want.Nums, twoSum(tc.Args.Nums, tc.Args.Target))
	}
}
