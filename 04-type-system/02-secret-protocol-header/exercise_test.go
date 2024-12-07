package secretprotocolheader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePublishFixHeader(t *testing.T) {
	assert.Equal(t, byte(68), createPublishFixHeader(false, false, false))
	assert.Equal(t, byte(69), createPublishFixHeader(false, false, true))
	assert.Equal(t, byte(70), createPublishFixHeader(false, true, false))
	assert.Equal(t, byte(71), createPublishFixHeader(false, true, true))
	assert.Equal(t, byte(84), createPublishFixHeader(true, false, false))
	assert.Equal(t, byte(85), createPublishFixHeader(true, false, true))
	assert.Equal(t, byte(86), createPublishFixHeader(true, true, false))
	assert.Equal(t, byte(87), createPublishFixHeader(true, true, true))
}
