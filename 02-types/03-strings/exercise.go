package strings

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func multilineString() string {
	return `some
multiline
string`
}

func stringLen(s string) int {
	return len(s)
}

func trimFirstChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[1:]
}

func trimLastChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[:len(s)-1]
}

func swapFirstChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return "A" + s[1:]
}

func swapLastChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[:len(s)-1] + "A"
}

func prependChar(s string) string {
	return "A" + s
}

func appendChar(s string) string {
	return s + "A"
}
