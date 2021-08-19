package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x%10 == 0 && x != 0 {
		return false
	}

	var revert int
	for x > revert {
		revert = revert*10 + x%10
		x /= 10
	}

	return (x == revert) || (x == revert/10)
}

func TestIsPalindrome(t *testing.T) {
	x := 121
	target := true
	require.Equal(t, target, isPalindrome(x))

	x = -121
	target = false
	require.Equal(t, target, isPalindrome(x))

	x = 10
	target = false
	require.Equal(t, target, isPalindrome(x))

	x = 5
	target = true
	require.Equal(t, target, isPalindrome(x))
}
