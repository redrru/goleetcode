package array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	seen := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	type (
		args struct {
			Nums []int
		}
		want struct {
			IsContains bool
		}
	)

	tests := []struct {
		Args args
		Want want
	}{
		{
			Args: args{
				Nums: []int{1, 2, 3, 1},
			},
			Want: want{
				IsContains: true,
			},
		},
		{
			Args: args{
				Nums: []int{1, 2, 3, 4},
			},
			Want: want{
				IsContains: false,
			},
		},
		{
			Args: args{
				Nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			Want: want{
				IsContains: true,
			},
		},
		{
			Args: args{
				Nums: []int{1},
			},
			Want: want{
				IsContains: false,
			},
		},
	}

	for _, tc := range tests {
		require.Equal(t, tc.Want.IsContains, containsDuplicate(tc.Args.Nums))
	}
}
