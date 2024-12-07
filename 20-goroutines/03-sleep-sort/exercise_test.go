package sleepSort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSleepSort(t *testing.T) {
	input := []uint{47, 1, 10, 23, 42, 7}
	result := []uint{1, 7, 10, 23, 42, 47}
	assert.Equal(t, result, sleepSort(input), "forward sleep-sort")
}
