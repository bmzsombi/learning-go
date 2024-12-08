package repaint

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// repaintColor returns the complementary of the color received as argument or an error for an unknown error.
func repaintColor(color string) (string, error) {
	// INSERT YOUR CODE HERE
	switch color {
	case "vermilion":
		return "teal", nil
	case "teal":
		return "vermilion", nil
	default:
		return "", fmt.Errorf("uknown color: %s", color)
	}
}
