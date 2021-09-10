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

func TestFindLanguageByName(t *testing.T) {
	assert.Equal(t, "Ruby", LanguagesById[326])
}

func TestGemfileLock(t *testing.T) {
	langs := LanguageForPath("Gemfile.lock")
	assert.Equal(t, 0, len(langs))
}

func TestUnableToDetect(t *testing.T) {
	langs := LanguageForPath("noideawhatthisis")
	assert.Equal(t, 0, len(langs))
}

func TestFields(t *testing.T) {
	ruby := Languages["Ruby"]
	assert.NotNil(t, ruby)
	assert.Equal(t, uint(326), ruby.ID)
	assert.Equal(t, "#701516", ruby.Color)
	assert.Equal(t, "source.ruby", ruby.TMScope)
}

func TestParsedLanguagesYml(t *testing.T) {
	assert.Equal(t, 598, len(Languages))
	assert.Equal(t, 1202, len(LanguagesByExtension))
	assert.Equal(t, 287, len(LanguagesByFileName))
}
