package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	require.False(t, BinarySearch(arr, 8))
	require.True(t, BinarySearch(arr, 7))
	require.True(t, BinarySearch(arr, 1))
	require.False(t, BinarySearch(arr, 0))
	require.True(t, BinarySearch(arr, 9))
}
