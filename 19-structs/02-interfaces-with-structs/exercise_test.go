package structsinterfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testArea(e Shape) float64 {
	return e.Area()
}

func testPerimeter(e Shape) float64 {
	return e.Perimeter()
}

func TestStructsInterfaces(t *testing.T) {
	s1 := NewCircle(5.0)

	assert.Equal(t, s1.Area(), testArea(s1))
	assert.Equal(t, s1.Perimeter(), testPerimeter(s1))

	s2 := NewRectangle(5.0,2.0)
	assert.Equal(t, s2.Area(), testArea(s2))
	assert.Equal(t, s2.Perimeter(), testPerimeter(s2))
}
