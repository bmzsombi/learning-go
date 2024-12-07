package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	assert.Equal(t, 2, amean(1.2, 2.4), "basic test 0")
	assert.Equal(t, 52, amean(3.67, 100.0), "basic test 1")
	assert.Equal(t, 89, amean(34.91, 144.02), "basic test 2")

	val, err := ameanString("3.67", "100.0")
	assert.NoError(t, err, "string parse test 1")
	assert.Equal(t, 52, val, "string test 1")

	val, err = ameanString("34.91", "144.02")
	assert.NoError(t, err, "string parse test 2")
	assert.Equal(t, 89, val, "string test 2")

	val, err = ameanString("dummmy", "100.0")
	assert.Error(t, err, "string parse test x")
	assert.Equal(t, 0, val, "string test 1")

	val, err = ameanString("1.0", "dummmy")
	assert.Error(t, err, "string parse test y")
	assert.Equal(t, 0, val, "string test 2")
}
