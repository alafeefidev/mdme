package mdme

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var defaultsIgnore = map[string]struct{}{
	".git":            {},
	".vscode":         {},
	"node_modules":    {},
	"__pycache__":     {},
	".DS_Store":       {},
	"vendor":          {},
	".terraform":      {},
	"dist":            {},
	"build":           {},
	".next":           {},
	".idea":           {},
	".svn":            {},
	".hg":             {},
	".env":            {},
	".env.local":      {},
	"requirements.py": {},
}

func ignore(name string, ignores map[string]struct{}) bool {
	_, ok := ignores[name]
	return ok
}

type IgnoreConfig struct {
	//TODO gitignore, and custom ignore file .mdignore
	defaults   map[string]struct{}
	hiddenDirs bool // Skip hidden directories, dotted.
}

var DefaultIgnoreConfig = &IgnoreConfig{
	defaults:   defaultsIgnore,
	hiddenDirs: true,
}

var PlainIgnoreConfig = &IgnoreConfig{}

func (im *IgnoreConfig) Ignore(path string, isDir bool) bool {
	name := filepath.Base(path)

	if ignore(name, im.defaults) {
		return true
	}

	if im.hiddenDirs && strings.HasPrefix(name, ".") {
		return true
	}

	return false
}

func ErrorMsg(err error) string {
	msg := err.Error()
	// if it contains a semicolon
	if i := strings.LastIndex(msg, ": "); i != -1 {
		return msg[i+2:]
	}
	// if not
	return msg
}

func IsTextFile(path string) (bool, error) {
	data := make([]byte, 512)
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	n, err := f.Read(data)
	if err != nil {
		return false, err
	}

	if slices.Contains(data[:n], 0) {
		return false, nil
	}
	return true, nil
}
