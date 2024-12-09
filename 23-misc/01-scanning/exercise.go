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
	scanner.Split(bufio.ScanWords) // Split input into words

	lowercaseCount := 0

	for scanner.Scan() {
		word := scanner.Text()
		// Tisztítsuk meg az írásjeleket a szóból
		cleanWord := clean(word)
		// Ellenőrizzük, hogy tisztán kisbetűs-e
		if isLowercaseWord(cleanWord) {
			lowercaseCount++
		}
	}

	return lowercaseCount
}

// Helper function to clean punctuation from a word
func clean(word string) string {
	var buffer bytes.Buffer
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// Helper function to check if a word is entirely lowercase
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
