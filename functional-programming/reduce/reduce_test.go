package fp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReduce_Example(t *testing.T) {
	res := Reduce(func(acc int, curr int, _ int) int { return acc + curr }, 0)([]int{1, 2, 3})
	require.Equal(t, 6, res, fmt.Sprintf("Reduce should call the specified callback function for all the elements in an array. The return value of the callback function should be the accumulated result, and is provided as an argument in the next call to the callback function. Expected: %v Received: %v", 6, res))
}
