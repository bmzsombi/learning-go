package filteringdata

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// filterData filters a slice based in an index slice.
func filterData(keys []string, indices []int) [10]string {
	// INSERT YOUR CODE HERE
	var result [10]string

	if len(keys) != len(indices) {
		return result
	}

	filteredKeys := []string{}
	for i, key := range keys {
		if indices[i] <= 5 {
			filteredKeys = append(filteredKeys, key)
		}
	}

	for i := 0; i < len(filteredKeys) && i < 10; i++ {
		result[i] = filteredKeys[i]
	}

	return result
}
