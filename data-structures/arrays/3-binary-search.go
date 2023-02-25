package arrays

func BinarySearch(data []int, value int) bool {
	start := 0
	end := len(data)

	for start <= end {
		mid := start + (end-start)/2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}
