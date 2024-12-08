package factorial

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func calcSum(n int) int {
	var sum int
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum
}
