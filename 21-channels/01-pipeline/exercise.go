package pipeline

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
// generator takes an array of integers and returns a channel that emits the integers
func generator(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

// adder takes a channel of integers, increments each by 3, and returns a channel of float32
func adder(in <-chan int) <-chan float32 {
	out := make(chan float32)
	go func() {
		for num := range in {
			out <- float32(num + 3)
		}
		close(out)
	}()
	return out
}

// collector reads numbers from a float32 channel and collects them into a slice
func collector(in <-chan float32) []float32 {
	var results []float32
	for num := range in {
		results = append(results, num)
	}
	return results
}
