package wordcount

import (
	"strings"
	"time"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func CountWords(words []string) map[string]int {
	if len(words) == 0 {
		return map[string]int{}
	}

	wordCountChan := make(map[string]int)
	wordChan := make(chan string)
	done := make(chan bool)
	defer close(wordChan)

	go func() {
		for i := range wordChan {
			wordCountChan[i]++
		}
		done <- true
	}()

	for _, str := range words {
		go func(s string) {
			wordsFields := strings.Fields(s)
			for _, word := range wordsFields {
				wordChan <- word
			}
		}(str)
	}
	time.Sleep(40 * time.Millisecond)
	return wordCountChan
}
