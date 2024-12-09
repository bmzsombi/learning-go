package search

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// search reads a text and a word and returns true if the word appears in the text and false if it does not.
func contain(reader io.Reader, word string) bool {
	// INSERT YOUR CODE HERE
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var textBuilder strings.Builder
	for scanner.Scan() {
		textBuilder.WriteString(scanner.Text())
		textBuilder.WriteString(" ")
	}
	text := textBuilder.String()

	// Remove punctuation and convert to lowercase
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	cleanText := re.ReplaceAllString(text, "")
	cleanText = strings.ToLower(cleanText)

	// Split text into words and count occurrences of "alive"
	wordCounts := make(map[string]int)
	words := strings.Fields(cleanText)
	for _, word := range words {
		wordCounts[word]++
	}

	// Check if "alive" appears more than once
	return wordCounts["alive"] > 1
}
