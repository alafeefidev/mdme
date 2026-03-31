package mdme

import (
	"path/filepath"
	"strings"

	gitignore "github.com/sabhiram/go-gitignore"
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
	defaults   map[string]struct{}
	hiddenDirs bool // Skip hidden directories, dotted.
	gitignore  *gitignore.GitIgnore
	mdignore   *gitignore.GitIgnore
}

func (im *IgnoreConfig) Ignore(path, rel string, isDir bool) bool {
	name := filepath.Base(path)

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
