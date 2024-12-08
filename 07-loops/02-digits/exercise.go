package digits

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func sumDigits(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
