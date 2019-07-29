package lingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectByExtension(t *testing.T) {
	lang := LanguageForPath("test.rb")
	assert.Equal(t, "Ruby", lang.Name)
}

func TestDetectByFileName(t *testing.T) {
	lang := LanguageForPath("Rakefile")
	assert.Equal(t, "Ruby", lang.Name)
}
