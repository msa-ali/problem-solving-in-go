package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapsackRecursive(t *testing.T) {
	assert.Equal(t, 55, FindMaxKnapsackProfit(30, []int{10, 20, 30}, []int{22, 33, 44}))
	assert.Equal(t, 15, FindMaxKnapsackProfit(5, []int{1, 2, 3, 5}, []int{10, 5, 4, 8}))
}

func TestKnapsackRecursiveMemo(t *testing.T) {
	assert.Equal(t, 55, FindMaxKnapsackProfitTopDown(30, []int{10, 20, 30}, []int{22, 33, 44}))
	assert.Equal(t, 15, FindMaxKnapsackProfitTopDown(5, []int{1, 2, 3, 5}, []int{10, 5, 4, 8}))
}

func TestKnapsackBottomUp(t *testing.T) {
	assert.Equal(t, 55, FindMaxKnapsackProfitBottomUp(30, []int{10, 20, 30}, []int{22, 33, 44}))
	assert.Equal(t, 15, FindMaxKnapsackProfitBottomUp(5, []int{1, 2, 3, 5}, []int{10, 5, 4, 8}))
}
