//go:build helloworldgoversion

package main

import (
	"io"
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldWithGoVersion(t *testing.T) {
	res, err := testHTTP(t, "/", "GET", "")
	assert.NoError(t, err, "GET")
	assert.Equal(t, http.StatusOK, res.StatusCode, "status code")

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "read response body")
	r := regexp.MustCompile("^Hello world from .* running Go version.*")
	assert.Regexp(t, r, string(body), "response")
}
