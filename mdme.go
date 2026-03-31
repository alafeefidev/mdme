package mdme

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path    string
	Content []byte
}

func ListFiles(root string, maxDepth, maxFiles int) ([]File, error) {
	var files []File

	conf := &IgnoreConfig{
		defaults:   defaultsIgnore,
		hiddenDirs: true,
		gitignore:  FromFile(filepath.Join(root, ".gitignore")),
		mdignore:   FromFile(filepath.Join(root, ".mdignore")),
	}

	rootDepth := strings.Count(filepath.Clean(root), string(os.PathSeparator))

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			slog.Warn("could not traverse into %s: %v\n", path, err)
			return nil // Skip instead of stopping
		}

		if len(files) >= maxFiles {
			return fmt.Errorf("too many files: limit is %d", maxFiles)
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			slog.Debug("Skipping", "path", rel)
			return nil
		}

		if maxDepth > 0 {
			currentDepth := strings.Count(filepath.Clean(path), string(os.PathSeparator)) - rootDepth
			if d.IsDir() && currentDepth >= maxDepth {
				slog.Debug("Skipping", "directory", rel)
				return filepath.SkipDir
			}
		}

		slog.Debug("Processing", "Path", rel)

		if conf.Ignore(path, rel, d.IsDir()) {
			// Skip directory and sub-directories and files
			slog.Debug("Skipping", "path", rel)
			if d.IsDir() {
				return filepath.SkipDir
			}
			// Skip file
			return nil
		}

		if !d.IsDir() {
			content, err := readTextFile(path)
			if err == nil && content != nil {
				files = append(files, File{
					Path:    path,
					Content: content,
				})
			}
		}
		return nil
	})

	return files, err
}

func IsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
