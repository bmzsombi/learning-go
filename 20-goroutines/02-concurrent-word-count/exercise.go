package wordcount

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func CountWords(words []string) map[string]int {
	if len(words) == 0 {
		return map[string]int{}
	}

	var returnmap map[string]int
	countchan := make(chan int)

	for i, val1 := range words {
		for j, val2 := range words {
			go func(word string, returnmap map[string]int) {
				defer close(countchan)
				if words[i] == words[j] {
					val1++
				}
			}(words[i], returnmap)
		}
	}
}
