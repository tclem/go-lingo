package lingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectByExtension(t *testing.T) {
	langs := LanguageForPath("test.rb")
	assert.Equal(t, 1, len(langs))
	assert.Equal(t, "Ruby", langs[0].Name)

	langs = LanguageForPath(".notanvalidext")
	assert.Equal(t, 0, len(langs))
}

func TestDetectByFileName(t *testing.T) {
	langs := LanguageForPath("Rakefile")
	assert.Equal(t, 1, len(langs))
	assert.Equal(t, "Ruby", langs[0].Name)
}

func TestGemfileLock(t *testing.T) {
	langs := LanguageForPath("Gemfile.lock")
	assert.Equal(t, 0, len(langs))
}

func TestUnableToDetect(t *testing.T) {
	langs := LanguageForPath("noideawhatthisis")
	assert.Equal(t, 0, len(langs))
}

func TestParsedLanguagesYml(t *testing.T) {
	assert.Equal(t, 555, len(Languages))
	assert.Equal(t, 1160, len(LanguagesByExtension))
	assert.Equal(t, 257, len(LanguagesByFileName))
}
