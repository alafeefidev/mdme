package mdme

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

func readTextFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if n == 0 {
		return nil, nil
	}
	if err != nil && err != io.EOF {
		return nil, err
	}

	buf = buf[:n]

	for _, b := range buf {
		if b == 0 {
			return nil, fmt.Errorf("file is binary: %s", path)
		}
	}

	contentType := http.DetectContentType(buf)
	if !strings.Contains(contentType, "text/") {
		return nil, fmt.Errorf("file is binary: %s", path)
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	return io.ReadAll(f)
}

func CreateMDFile(name, path, MD string) error {
	mdName := name
	if !strings.HasSuffix(mdName, ".md") {
		mdName = mdName + ".md"
	}
	fPath := filepath.Join(path, mdName)
	return os.WriteFile(fPath, []byte(MD), 0644)
}

func IsHomeDir(absPath string) bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	return strings.EqualFold(absPath, filepath.Clean(home))
}

