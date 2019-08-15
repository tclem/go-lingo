// Package go-lingo is for basic programming language detection of source code files.
//go:generate go run ./cmd/lingo/lingo.go
package lingo

import (
	"path/filepath"
)

// LanguageForPath returns the programming language (if any) for a filepath.
func LanguageForPath(path string) *Language {
	ext := filepath.Ext(path)
	lang, ok := Languages[LanguagesByExtension[ext]]
	if ok {
		return &lang
	}

	name := filepath.Base(path)
	lang, ok = Languages[LanguagesByFileName[name]]
	if ok {
		return &lang
	}

	return nil
}
