package arrays

// Given an array as an input argument,
// write a method that will return the sum of all the integer elements in the array.

func SumArr(data []int) int {
	total := 0
	for i := 0; i < len(data); i++ {
		total += data[i]
	}
	return total
}
