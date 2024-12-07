package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcFactorialOrSumOrAbs(t *testing.T) {
	
	
	assert.Equal(t, 55, calcSum(10))
	assert.Equal(t, 10, calcSum(4))
	assert.Equal(t, 1, calcSum(1))
	
	
}
