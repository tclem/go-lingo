package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type Language struct {
	ID         uint `yaml:"language_id"`
	Name       string
	TMScope    string   `yaml:"tm_scope"`
	Extensions []string `yaml:"extensions"`
	Filenames  []string `yaml:"filenames"`
}

// NOTE: Due to Go's random iteration of maps, generating will always change languages.go
func main() {
	log.SetFlags(0)
	log.SetPrefix("lingo: ")

	content, err := ioutil.ReadFile("../languages.yml")
	if err != nil {
		panic(err)
	}
	languages := map[string]Language{}
	err = yaml.Unmarshal([]byte(content), &languages)
	if err != nil {
		panic(err)
	}

	sortedLanguageNames := make([]string, 0, len(languages))
	for name := range languages {
		sortedLanguageNames = append(sortedLanguageNames, name)
	}
	sort.Strings(sortedLanguageNames)

	g := Generator{}
	g.Printf("// Code generated by \"lingo\"; DO NOT EDIT.\n")
	g.Printf("\n")
	g.Printf("package lingo\n")
	g.Printf("\n")
	g.Printf("type Language struct {\n")
	g.Printf("\tID uint\n")
	g.Printf("\tName string\n")
	g.Printf("\tTMScope string\n")
	g.Printf("\tExtensions []string\n")
	g.Printf("\tFilenames []string\n")
	g.Printf("}\n")
	g.Printf("\n")

	g.Printf("var (\n")
	languagesByExtension := map[string][]Language{}
	languagesByFileName := map[string][]Language{}
	languagesById := map[uint]string{}
	g.Printf("\tLanguages = map[string]Language{\n")
	for _, name := range sortedLanguageNames {
		v := languages[name]
		v.Name = name
		g.Printf("\t\t\"%s\":", name)
		g.printLanguage(&v)
		g.Printf(",\n")

		for _, e := range v.Extensions {
			x := languagesByExtension[e]
			x = append(x, v)
			languagesByExtension[e] = x
		}

		for _, e := range v.Filenames {
			x := languagesByFileName[e]
			x = append(x, v)
			languagesByFileName[e] = x
		}

		languagesById[v.ID] = name
	}
	g.Printf("}\n")

	sortedByExtension := mapSort(languagesByExtension)

	// Languages by extension
	g.Printf("\tLanguagesByExtension = map[string][]string{\n")
	for _, ext := range sortedByExtension {
		langs := languagesByExtension[ext]
		var names []string
		for _, l := range langs {
			names = append(names, fmt.Sprintf(`"%s"`, l.Name))
		}
		g.Printf("\t\t\"%s\": []string{%s},\n", ext, strings.Join(names, ", "))
	}
	g.Printf("\t}\n")

	sortedByFileName := mapSort(languagesByFileName)

	// Languages by filename
	g.Printf("\tLanguagesByFileName = map[string][]string{\n")
	for _, name := range sortedByFileName {
		langs := languagesByFileName[name]
		var names []string
		for _, l := range langs {
			names = append(names, fmt.Sprintf(`"%s"`, l.Name))
		}
		g.Printf("\t\t\"%s\": []string{%s},\n", name, strings.Join(names, ", "))
	}
	g.Printf("\t}\n")

	// Languages by id
	sortedById := make([]uint, 0, len(languagesById))
	for id := range languagesById {
		sortedById = append(sortedById, id)
	}
	sort.Slice(sortedById, func(i, j int) bool { return sortedById[i] < sortedById[j] })
	g.Printf("\tLanguagesById = map[uint]string{\n")
	for _, id := range sortedById {
		name := fmt.Sprintf(`"%s"`, languagesById[id])
		g.Printf("\t\t%d: %s,\n", id, name)
	}
	g.Printf("\t}\n")

	// end of `var` declaration
	g.Printf(")\n")

	// Format the output.
	src := g.format()

	err = ioutil.WriteFile("../lingo/languages.go", src, 0644)
	if err != nil {
		panic(err)
	}
}

func mapSort(toSort map[string][]Language) []string {
	sorted := make([]string, 0, len(toSort))
	for key := range toSort {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)
	return sorted
}

func (g *Generator) printLanguage(language *Language) {
	var extensions []string
	for _, e := range language.Extensions {
		extensions = append(extensions, fmt.Sprintf(`"%s"`, e))
	}
	exts := strings.Join(extensions, ", ")
	g.Printf("Language{ID: %d, Name:\"%s\", TMScope:\"%s\", Extensions: []string{%s} }", language.ID, language.Name, language.TMScope, exts)
}

type Generator struct {
	buf bytes.Buffer // Accumulated output.
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}
