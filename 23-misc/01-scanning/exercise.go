package scanning

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// counter reads a text and returns the counted values.
func counter(reader io.Reader) int {
	// INSERT YOUR CODE HERE
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	lowercaseCount := 0

	for scanner.Scan() {
		word := scanner.Text()
		cleanWord := clean(word)
		if isLowercaseWord(cleanWord) {
			lowercaseCount++
		}
	}

	return lowercaseCount
}

func clean(word string) string {
	var buffer bytes.Buffer
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

func isLowercaseWord(word string) bool {
	if word == "" {
		return false
	}
	for _, r := range word {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
