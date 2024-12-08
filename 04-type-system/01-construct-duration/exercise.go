package constructduration

import "time"

//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// constructTime constructs a `Time` instant based on its two arguments (arg1, arg2)
func constructDuration(arg1 int, arg2 int) time.Duration {
	// INSERT YOUR CODE HERE
	return time.Duration(arg1)*time.Hour + time.Duration(arg2)*time.Minute
}
