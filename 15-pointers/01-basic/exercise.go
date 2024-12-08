package pointerbasic

// DO NOT REMOVE THIS COMMENT
//
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate
func retrieveValue(pointer *bool) bool {
	// INSERT YOUR CODE HERE
	return *pointer
}
