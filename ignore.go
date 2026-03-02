package mdme

import (
	"path/filepath"
	"strings"

	gignore "github.com/sabhiram/go-gitignore"
)

var defaultsIgnore = map[string]struct{}{
	".git":             {},
	".vscode":          {},
	"node_modules":     {},
	"__pycache__":      {},
	".DS_Store":        {},
	"vendor":           {},
	".terraform":       {},
	"dist":             {},
	"build":            {},
	".next":            {},
	".idea":            {},
	".svn":             {},
	".hg":              {},
	".env":             {},
	".env.local":       {},
	"requirements.txt": {},
	"LICENSE":          {},
	".gitignore":       {},
	".mdignore":        {},
}

type IgnoreConfig struct {
	//TODO gitignore, and custom ignore file .mdignore
	defaults   map[string]struct{}
	hiddenDirs bool // Skip hidden directories, dotted.
	gitignore  *gignore.GitIgnore
	mdignore   *gignore.GitIgnore
}

func (im *IgnoreConfig) Ignore(rel string, isDir bool) bool {
	name := filepath.Base(rel)

	if _, ok := im.defaults[name]; ok {
		return true
	}

	// Only ignore dotted directories, not files to not ignore .gitignore, .mdignore, etc
	if im.hiddenDirs && strings.HasPrefix(name, ".") && isDir {
		return true
	}

	match := rel
	if isDir {
		match = rel + "/"
	}

	if im.gitignore != nil && im.gitignore.MatchesPath(match) {
		return true
	}
	if im.mdignore != nil && im.mdignore.MatchesPath(match) {
		return true
	}

	return false
}
