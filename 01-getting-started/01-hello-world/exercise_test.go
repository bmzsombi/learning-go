package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "Hallo Welt!", helloWorld(), "hello world in German")
}
