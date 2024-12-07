package pointerbasic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveValue(t *testing.T) {
	

	

	
	var x bool
	x = true
	assert.Equal(t, x, retrieveValue(&x))
	x = false
	assert.Equal(t, x, retrieveValue(&x))
	
}
