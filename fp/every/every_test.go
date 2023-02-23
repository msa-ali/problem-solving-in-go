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
	everyIsOdd := Every(isOdd)
	arr := []int{1, 3, 5, 7, 9}
	result := everyIsOdd(arr)
	require.True(t, result)

	test := Every(isGreaterThanOne)
	result = test(arr)
	require.False(t, result)
}
