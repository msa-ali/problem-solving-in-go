package concurrency_pattern

func process(v int) int {
	return v
}

func ProcessChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func(i int) {
			results <- process(i)
		}(i)
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}
