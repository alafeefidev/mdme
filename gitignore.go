package mdme

import (
	gitignore "github.com/sabhiram/go-gitignore"
)

func FromFile(path string) *gitignore.GitIgnore {
	// Skip error if not found
	gi, _ := gitignore.CompileIgnoreFile(path)
	return gi
}
