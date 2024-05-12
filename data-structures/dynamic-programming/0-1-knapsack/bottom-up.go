package dynamicprogramming

func FindMaxKnapsackProfitBottomUp(capacity int, weights []int, values []int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= capacity; j++ {
			if weights[i-1] <= j {
				dp[i][j] = max(
					dp[i-1][j],
					values[i-1]+dp[i-1][j-weights[i-1]],
				)
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}
	}
	return dp[n][capacity]
}
