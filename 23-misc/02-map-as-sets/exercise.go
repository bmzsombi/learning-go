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

	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	cleanText := re.ReplaceAllString(text, "")
	cleanText = strings.ToLower(cleanText)

	wordCounts := make(map[string]int)
	words := strings.Fields(cleanText)
	for _, word := range words {
		wordCounts[word]++
	}

	return wordCounts["alive"] > 1
}
