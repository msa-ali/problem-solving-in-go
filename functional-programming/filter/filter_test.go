package fp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func isOdd(n int, _ int) bool {
	return n%2 != 0
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filterOdd := Filter(isOdd)
	expected := []int{1, 3, 5, 7, 9}
	result := filterOdd(arr)
	require.True(t, reflect.DeepEqual(result, expected))
}
