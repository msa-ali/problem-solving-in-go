package dynamicprogramming

/*
You are given 𝑛 items whose weights and values are known, as well as a knapsack to carry these items.
The knapsack cannot carry more than a certain maximum weight, known as its capacity.

You need to maximize the total value of the items in your knapsack, while ensuring that
the sum of the weights of the selected items does not exceed the capacity of the knapsack.

If there is no combination of weights whose sum is within the capacity constraint, return 0.
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxKnapsackProfitHelper(capacity int, weights []int, values []int, n int) int {
	if n == 0 || capacity <= 0 {
		return 0
	}
	if weights[n-1] <= capacity {
		return max(
			values[n-1]+findMaxKnapsackProfitHelper(capacity-weights[n-1], weights, values, n-1),
			findMaxKnapsackProfitHelper(capacity, weights, values, n-1),
		)
	}
	return findMaxKnapsackProfitHelper(capacity, weights, values, n-1)
}

func FindMaxKnapsackProfit(capacity int, weights []int, values []int) int {
	return findMaxKnapsackProfitHelper(capacity, weights, values, len(weights))
}
