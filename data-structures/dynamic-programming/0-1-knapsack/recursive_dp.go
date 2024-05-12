package dynamicprogramming

/*
You are given 𝑛 items whose weights and values are known, as well as a knapsack to carry these items.
The knapsack cannot carry more than a certain maximum weight, known as its capacity.

You need to maximize the total value of the items in your knapsack, while ensuring that
the sum of the weights of the selected items does not exceed the capacity of the knapsack.

If there is no combination of weights whose sum is within the capacity constraint, return 0.
*/

func findMaxKnapsackProfitTopDownHelper(capacity int, weights []int, values []int, n int, dp [][]int) int {
	if n == 0 || capacity <= 0 {
		return 0
	}
	if dp[n-1][capacity] != -1 {
		return dp[n-1][capacity]
	}
	if weights[n-1] <= capacity {
		dp[n-1][capacity] = max(
			values[n-1]+findMaxKnapsackProfitTopDownHelper(capacity-weights[n-1], weights, values, n-1, dp),
			findMaxKnapsackProfitTopDownHelper(capacity, weights, values, n-1, dp),
		)
	} else {
		dp[n-1][capacity] = findMaxKnapsackProfitTopDownHelper(capacity, weights, values, n-1, dp)
	}
	return dp[n-1][capacity]
}

func FindMaxKnapsackProfitTopDown(capacity int, weights []int, values []int) int {
	return findMaxKnapsackProfitTopDownHelper(capacity, weights, values, len(weights), createDPArray(len(weights), capacity+1))
}

func createDPArray(m, n int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return dp
}
