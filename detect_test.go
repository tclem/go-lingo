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

func TestUnableToDetect(t *testing.T) {
	lang := LanguageForPath("noideawhatthisis")
	assert.Nil(t, lang)
}

func TestParsedLanguagesYml(t *testing.T) {
	assert.Equal(t, 519, len(Languages))
	assert.Equal(t, 1117, len(LanguagesByExtension))
	assert.Equal(t, 235, len(LanguagesByFileName))
}
