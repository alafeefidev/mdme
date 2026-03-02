package mdme

import (
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

type File struct {
	Path    string
	Content []byte
}

func ListFiles(root string) ([]File, error) {
	var files []File

	conf := &IgnoreConfig{
		defaults:   defaultsIgnore,
		hiddenDirs: true,
		gitignore:  FromFile(filepath.Join(root, ".gitignore")),
		mdignore:   FromFile(filepath.Join(root, ".mdignore")),
	}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			slog.Warn("could not traverse into %s: %v\n", path, err)
			return nil // Skip instead of stopping
		}

		rel, err := filepath.Rel(root, path)

		slog.Debug("Processing", "Path", rel)
		
		if err != nil {
			slog.Debug("Skipping", "path", rel)
			return nil
		}

		if conf.Ignore(path, rel, d.IsDir()) {
			// Skip directory and sub-directories and files
			slog.Debug("Skipping", "path", rel)
			if d.IsDir() {
				return filepath.SkipDir
			}
			// Skip file
			return nil
		}

		// Check if it is a proper text file and not a binary
		if !d.IsDir() {
			if data, _ := IsTextFile(path); data != nil {
				files = append(files, File{
					Path:    path,
					Content: data,
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
