package fp

func Map[T any, U any](fn func(item T, index int) U) func([]T) []U {
	return func(arr []T) []U {
		size := len(arr)
		result := make([]U, 0, size)
		for i := 0; i < size; i++ {
			result = append(result, fn(arr[i], i))
		}
		return result
	}
}
