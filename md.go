package mdme

import (
	"errors"
	"path/filepath"
	"strings"
)

func ToMD(files []File, root string) (string, error) {
	var sb strings.Builder

	for i, file := range files {
		rel, _ := filepath.Rel(root, file.Path)
		content := file.Content
		if len(content) == 0 {
			//TODO log if content empty
			continue
		}

		ext := strings.TrimPrefix(filepath.Ext(file.Path), ".")
		if ext == "" {
			ext = "text"
		}

		if i > 0 {
			sb.WriteString("\n\n")
		}

		sb.WriteString("```")
		sb.WriteString(ext)
		sb.WriteString(" - ")
		sb.WriteString(rel)
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
