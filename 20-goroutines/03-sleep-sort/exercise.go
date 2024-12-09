package sleepSort

import "time"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
// sleepSort returns the input uint-slice sorted in the forward order.
func sleepSort(input []uint) []uint {
	resultChan := make(chan uint, len(input))
	results := make([]uint, 0, len(input))

	for _, num := range input {
		go func(n uint) {
			time.Sleep(time.Duration(n*10) * time.Millisecond)
			resultChan <- n
		}(num)
	}

	for range input {
		results = append(results, <-resultChan)
	}

	return results
}
