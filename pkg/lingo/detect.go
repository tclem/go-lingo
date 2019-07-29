// Package lingo is for basic programming language detection of source code files.
//go:generate github.com/github/lingo/cmd/lingo
package lingo

import (
	"path/filepath"
)

// LanguageForPath returns the programming language (if any) for a filepath.
func LanguageForPath(path string) *Language {
	ext := filepath.Ext(path)
	langs := LanguagesByExtension[ext]
	if langs != nil && len(langs) > 0 {
		return &langs[0]
	}

	name := filepath.Base(path)
	langs = LanguagesByFileName[name]
	if langs != nil && len(langs) > 0 {
		return &langs[0]
	}

	return nil
}
