package narithmetic

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// nArithmetic returns the result of an arithmetic operation over "n" elements.
func nArithmetic(elems [10]int) int {
	// INSERT YOUR CODE HERE
	return elems[0] - elems[1] - elems[2] - elems[3] - elems[4] - elems[5] - elems[6] - elems[7] - elems[8] - elems[9]
}
