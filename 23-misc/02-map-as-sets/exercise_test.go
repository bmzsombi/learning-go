package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestContain(t *testing.T) {
	var stdin bytes.Buffer
	stdin.WriteString("The upper realm has plenary powers. He lived a long time ago, but he is still alive. Apollonius of Tyana, writing as Hermes Trismegistos, said, That which is above is that which is below. By this he meant to tell us that our universe is a hologram, but he lacked the term.")

	assert.Equal(t, false, contain(&stdin, "alive"), "The result should be: false")
}
