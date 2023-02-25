package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxSubArraySum(t *testing.T) {
	data := []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	expected := 9
	result := MaxSubArraySum(data)
	require.Equal(t, expected, result)
}
