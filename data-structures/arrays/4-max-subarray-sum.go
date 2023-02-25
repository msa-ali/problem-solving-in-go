package arrays

func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSum := 0
	currentMax := 0
	for i := 0; i < size; i++ {
		currentMax = currentMax + data[i] // add current value
		// if current max is negative, reset to 0
		if currentMax < 0 {
			currentMax = 0
		}
		// if current max is greater than maxSum, update maxSum
		if maxSum < currentMax {
			maxSum = currentMax
		}
	}
	return maxSum
}
