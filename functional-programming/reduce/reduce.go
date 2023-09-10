package fp

func Reduce[T any, V any](reducer func(acc V, current T, index int) V, init V) func(arr []T) V {
	return func(arr []T) V {
		len := len(arr)
		result := init
		for i := 0; i < len; i++ {
			result = reducer(result, arr[i], i)
		}
		return result
	}
}
