package readsecretregister

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseChannelControlRegister(t *testing.T) {
	a,b,c,d := parseChannelControlRegister(1442060820)
	assert.Equal(t, []byte{0x1A, 0xF4, 0x55, 0x14}, []byte{a,b,c,d})

	a,b,c,d = parseChannelControlRegister(0xdeadbeef)
	assert.Equal(t, []byte{0xBE, 0xAD, 0xDE, 0xEF}, []byte{a,b,c,d})

	a,b,c,d = parseChannelControlRegister(0x01234567)
	assert.Equal(t, []byte{0x45, 0x23, 0x01, 0x67}, []byte{a,b,c,d})
}
