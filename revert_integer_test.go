package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	num := int32(x)

	var res, prev int32
	for num != 0 {
		res, prev = res*10+num%10, res
		if (res / 10) != prev {
			return 0
		}

		num = num / 10
	}

	return int(res)
}

func TestRevertInteger(t *testing.T) {
	x := 123
	target := 321
	require.Equal(t, target, reverse(x))

	x = -123
	target = -321
	require.Equal(t, target, reverse(x))

	x = 1200
	target = 21
	require.Equal(t, target, reverse(x))

	x = 0
	target = 0
	require.Equal(t, target, reverse(x))

	x = 1534236469
	target = 0
	require.Equal(t, target, reverse(x))
}
