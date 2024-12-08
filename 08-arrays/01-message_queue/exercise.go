package messagequeue

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// messagequeue returns the an array constructed from the arguments
func messageQueue(a, b, c string) [3]string {
	// INSERT YOUR CODE HERE
	var message [3]string
	message[0] = a
	message[1] = c
	message[2] = b
	return message
}
