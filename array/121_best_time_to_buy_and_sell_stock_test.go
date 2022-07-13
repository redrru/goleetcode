package array

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	profit := 0
	minPrice := math.MaxInt
	for _, price := range prices {
		minPrice = min(minPrice, price)
		profit = max(profit, price-minPrice)
	}

	return profit
}

func max(r, l int) int {
	if r > l {
		return r
	}
	return l
}

func min(r, l int) int {
	if r < l {
		return r
	}
	return l
}

func TestMaxProfit(t *testing.T) {
	type (
		args struct {
			Prices []int
		}
		want struct {
			Profit int
		}
	)

	tests := []struct {
		Args args
		Want want
	}{
		{
			Args: args{
				Prices: []int{7, 1, 5, 3, 6, 4},
			},
			Want: want{
				Profit: 5,
			},
		},
		{
			Args: args{
				Prices: []int{7, 6, 4, 3, 1},
			},
			Want: want{
				Profit: 0,
			},
		},
		{
			Args: args{
				Prices: []int{},
			},
			Want: want{
				Profit: 0,
			},
		},
		{
			Args: args{
				Prices: []int{1},
			},
			Want: want{
				Profit: 0,
			},
		},
		{
			Args: args{
				Prices: []int{1, 2},
			},
			Want: want{
				Profit: 1,
			},
		},
		{
			Args: args{
				Prices: []int{2, 1},
			},
			Want: want{
				Profit: 0,
			},
		},
	}

	for _, tc := range tests {
		require.Equal(t, tc.Want.Profit, maxProfit(tc.Args.Prices))
	}
}
