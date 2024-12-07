package pathsplit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitPath(t *testing.T) {
	assert.Equal(t, "jquery.js", splitPath("static/js/jquery.js"))
	assert.Equal(t, "multi_langs.js", splitPath("multi_langs.js"))
	assert.Equal(t, "style.css", splitPath("css/style.css"))
	assert.Equal(t, "flags.css", splitPath("css/flags.css"))
	assert.Equal(t, "favicon.ico", splitPath("image/favicon.ico"))
}
