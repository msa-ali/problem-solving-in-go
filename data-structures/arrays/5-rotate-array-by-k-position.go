package arrays

func reverse(arr []int, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}

func RotateArray(data []int, k int) {
	// reverse array from index 0 to k-1
	reverse(data, 0, k)

	// reverse  from k to n-1
	reverse(data, k, len(data))

	// reverse the whole array
	reverse(data, 0, len(data))
}
