package calculator

import (
	"math"
	"strconv"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func amean(x, y float64) int {
	mean := (x + y) / 2
	return int(math.Round(mean))
}

func ameanString(x, y string) (int, error) {
	xf, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return 0, err
	}

	xy, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return 0, err
	}

	mean := (xf + xy) / 2
	return int(math.Round(mean)), err
}
