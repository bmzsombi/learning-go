package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := []float32{4, 5, 6, 7, 8, 9, 10, 11}
	assert.Equal(t, result, collector(adder(generator(input))))
}
