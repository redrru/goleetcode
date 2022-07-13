package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/majority-element/

func majorityElementSort(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[len(nums)/2]
}

func TestMajorityElement(t *testing.T) {
	type (
		args struct {
			Nums []int
		}
		want struct {
			Majority int
		}
	)

	tests := []struct {
		Args args
		Want want
	}{
		{
			Args: args{
				Nums: []int{3, 2, 3},
			},
			Want: want{
				Majority: 3,
			},
		},
		{
			Args: args{
				Nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			Want: want{
				Majority: 2,
			},
		},
		{
			Args: args{
				Nums: []int{2},
			},
			Want: want{
				Majority: 2,
			},
		},
	}

	for _, tc := range tests {
		require.Equal(t, tc.Want.Majority, majorityElementSort(tc.Args.Nums))
	}
}
