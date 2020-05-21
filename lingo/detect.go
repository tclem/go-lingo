// Package go-lingo is for basic programming language detection of source code files.
//go:generate go run ../cmd/lingo/lingo.go
package lingo

import (
	"path/filepath"
)

// LanguageForPath returns the programming languages (if any) for a filepath.
func LanguageForPath(path string) []Language {
	ext := filepath.Ext(path)
	names := LanguagesByExtension[ext]
	if len(names) > 0 {
		return lookup(names)
	}

	name := filepath.Base(path)
	names = LanguagesByFileName[name]
	return lookup(names)
}

func lookup(names []string) []Language {
	var langs []Language
	for _, n := range names {
		langs = append(langs, Languages[n])
	}
	return langs
}
