package mdme

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)


func ListFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("could not traverse into %s: %v\n", path, err)
			return nil // Skip instead of stopping
		}

		if DefaultIgnoreConfig.Ignore(path, d.IsDir()) {
			// Skip directory and sub-directories and files
			if d.IsDir() {
				return filepath.SkipDir
			}
			// Skip file
			return nil
		}
		if ok, err := IsTextFile(path); ok && !d.IsDir(){
			files = append(files, path)
		
		} else if err != nil {
			return nil
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func IsDir(path string) (bool, error){
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}