package greedytechniques

import (
	"sort"
)

// https://leetcode.com/problems/two-city-scheduling/description/

func twoCityScheduling(costs [][]int) int {
	sort.SliceStable(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})
	n := len(costs)
	totalCost := 0
	for i := 0; i < n/2; i++ {
		totalCost += costs[i][0] + costs[n-i-1][1]
	}
	return totalCost
}
