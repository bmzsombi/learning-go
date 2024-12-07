package digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigits(t *testing.T) {
	
	
		assert.Equal(t, 7, sumDigits(7))
	
		assert.Equal(t, 3, sumDigits(111))
	
		assert.Equal(t, 1, sumDigits(1000))
	
		assert.Equal(t, 45, sumDigits(1307674368000))
	
		assert.Equal(t, 46, sumDigits(130767436801))
	
	
	
}
