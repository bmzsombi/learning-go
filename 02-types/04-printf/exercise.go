package printer

import (
	"fmt"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func printBool(b bool) string {
	return fmt.Sprintf("variable of type boolean and value ", b)
}

func printInt(i int) string {
	return fmt.Sprint("variable of type integer and value ", i)
}

func printHex(i int) string {
	return fmt.Sprintf("variable of type integer in hexadecimal form and value %X", i)
}

func printFloat(f float64) string {
	return fmt.Sprintf("ariable of type float and value %.2f", f)
}

func printString(s string) string {
	return fmt.Sprintf("variable of type string and value \"%s\"", s)
}

func concatStrings(a, b string) string {
	return a + b
}

func printConcatStrings(a, b string) string {
	return fmt.Sprint(printString(a + b))
}
