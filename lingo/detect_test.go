package lingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectByExtension(t *testing.T) {
	lang := LanguageForPath("test.rb")
	assert.Equal(t, "Ruby", lang.Name)

	lang = LanguageForPath(".notanvalidext")
	assert.Nil(t, lang)
}

func TestDetectByFileName(t *testing.T) {
	lang := LanguageForPath("Rakefile")
	assert.Equal(t, "Ruby", lang.Name)
}

func TestGemfileLock(t *testing.T) {
	lang := LanguageForPath("Gemfile.lock")
	assert.Nil(t, lang)
}

func TestUnableToDetect(t *testing.T) {
	lang := LanguageForPath("noideawhatthisis")
	assert.Nil(t, lang)
}

func TestParsedLanguagesYml(t *testing.T) {
	assert.Equal(t, 519, len(Languages))
	assert.Equal(t, 1117, len(LanguagesByExtension))
	assert.Equal(t, 234, len(LanguagesByFileName))
}
