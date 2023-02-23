package fp

func Filter[T any](fn func(item T, index int) bool) func([]T) []T {
	return func(arr []T) []T {
		len := len(arr)
		result := make([]T, 0, len)
		for i := 0; i < len; i++ {
			if fn(arr[i], i) {
				result = append(result, arr[i])
			}
		}
		return result
	}
}
