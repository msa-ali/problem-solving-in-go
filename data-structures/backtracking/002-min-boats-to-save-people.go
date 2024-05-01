package backtracking

import "sort"

func rescueBoats(people []int, limit int) int {
	// slices.Sort(people)
	sort.Ints(people)
	numOfBoats, left, right := 0, 0, len(people)-1
	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		numOfBoats++
	}
	return numOfBoats
}
