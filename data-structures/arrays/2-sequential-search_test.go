package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSequentialSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := SequentialSearch(arr, 3)
	require.True(t, result)

	result = SequentialSearch(arr, 10)
	require.False(t, result)
}
