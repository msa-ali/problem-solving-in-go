package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumArr(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := SumArr(data)
	require.Equal(t, sum, 55)
	require.Equal(t, SumArr([]int{}), 0)
}
