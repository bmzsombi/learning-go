package constructduration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructDuration(t *testing.T) {
	assert.Equal(t, `1h2m0s`, constructDuration(1, 2).String())
	assert.Equal(t, `4h7m0s`, constructDuration(4, 7).String())
	assert.Equal(t, `12h23m0s`, constructDuration(12, 23).String())
}
