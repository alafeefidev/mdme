package mdme

import (
	"strings"

	gignore "github.com/sabhiram/go-gitignore"
)

func FromFile(path string) *gignore.GitIgnore {
	// Skip error if not found
	gi, _ := gignore.CompileIgnoreFile(path)
	return gi
}

func fromBytes(data []byte) *gignore.GitIgnore {
	lines := strings.Split(string(data), "\n")
	return gignore.CompileIgnoreLines(lines...)
}