package logicalops

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func inverse(b bool) bool {
	// INSERT YOUR CODE HERE
	return !b
}

func and(x, y bool) bool {
	return x && y
}

func deMorgan(a, b bool) bool {
	return and(inverse(a), inverse(b))
}
