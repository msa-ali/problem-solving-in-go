package backtracking

func sum(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func gasStationJourney(gas []int, cost []int) int {
	if sum(cost) > sum(gas) {
		return -1
	}
	total := 0
	start := 0

	for i := range len(gas) {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			start = i + 1
		}
	}
	return start
}
