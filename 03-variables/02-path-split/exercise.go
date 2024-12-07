package pathsplit

import "path"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
// splitPath returns the file component of a file path.
func splitPath(fullPath string) string {
	_, file := path.Split(fullPath)
	return file
}
