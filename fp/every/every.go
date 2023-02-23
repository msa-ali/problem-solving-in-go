package fp

func Every[T any](fn func(item T, index int) bool) func([]T) bool {
	return func(arr []T) bool {
		len := len(arr)
		for i := 0; i < len; i++ {
			if !fn(arr[i], i) {
				return false
			}
		}
		return true
	}
}
