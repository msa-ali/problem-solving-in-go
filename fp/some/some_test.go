package fp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func isOdd(n int, _ int) bool {
	return n%2 != 0
}

func isGreaterThanOne(n int, _ int) bool {
	return n > 1
}

func TestFilter(t *testing.T) {
	someIsOdd := Some(isOdd)
	arr := []int{2, 4, 6, 8, 10}
	result := someIsOdd(arr)
	require.False(t, result)

	test := Some(isGreaterThanOne)
	result = test(arr)
	require.True(t, result)
}
