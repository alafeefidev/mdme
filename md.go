package mdme

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ToMD(files []string) (string, error) {
	var sb strings.Builder

	for i, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			//TODO log just skip the file if an error occured
			continue
		}
		if len(content) == 0 {
			//TODO log if content empty
			continue
		}

		ext := strings.TrimPrefix(filepath.Ext(file), ".")
		if ext == "" {
			ext = "text"
		}

		if i > 0 {
			sb.WriteString("\n\n")
		}

		sb.WriteString("```")
		sb.WriteString(ext)
		sb.WriteString("\n")
		sb.Write(content)
		if content[len(content)-1] != '\n' {
			sb.WriteByte('\n')
		}
		sb.WriteString("```")
	}

	if sb.Len() == 0 {
		return "", errors.New("nothing to print")
	}

	return sb.String(), nil
}
