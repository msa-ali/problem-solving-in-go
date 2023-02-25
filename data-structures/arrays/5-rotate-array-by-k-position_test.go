package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 2
	expected := []int{3, 4, 5, 6, 1, 2}
	RotateArray(arr, k)
	require.Equal(t, expected, arr)
}
