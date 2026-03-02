package mdme

import (
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func ErrorMsg(err error) string {
	msg := err.Error()
	// if it contains a semicolon
	if i := strings.LastIndex(msg, ": "); i != -1 {
		return msg[i+2:]
	}
	// if not
	return msg
}

func IsTextFile(path string) ([]byte, error) {
	data := make([]byte, 512)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	n, err := f.Read(data)
	if n == 0 {
		return nil, nil
	}
	if err != nil && err != io.EOF {
		return nil, err
	}

	if slices.Contains(data[:n], 0) {
		return nil, nil
	}
	return data, nil
}

func IsHomeDir(absPath string) bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	return strings.EqualFold(absPath, filepath.Clean(home))
}