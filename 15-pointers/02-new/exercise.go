package pointernew

// DO NOT REMOVE THIS COMMENT
//
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate
func newValue() *bool {
	// INSERT YOUR CODE HERE
	value := new(bool)
	*value = false
	return value
}
