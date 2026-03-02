package mdme

//=====================================
//=========TODO IN DEVELOPMENT=========
//=====================================

// import (
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"unicode/utf8"
// )

// type GitIgnoreList struct {
// 	files       []string
// 	directories []string
// 	paths       []string
// 	wildcards   []string
// }

// //TODO another goroutine that also loop the files and check the wildcards

// func (gi *GitIgnoreList) GitIgnoreInit(path string) error {
// 	// path is the gitignore (or custom ignore) file path

// 	dir := filepath.Dir(path)
// 	igData, err := os.ReadFile(path)
// 	if err != nil {
// 		return err
// 	}

// 	lines := strings.Split(strings.TrimSpace(string(igData)), "\n")
// 	for i, line := range lines {
// 		lines[i] = strings.TrimSpace(line)

// 		// Comment
// 		if strings.HasPrefix(line, "#") {
// 			continue
// 		}

// 		if strings.Contains(line, "*") {

// 		}

// 		//TODO remove
// 		if !strings.Contains(line, "*") && !strings.Contains(line, "?") && !strings.Contains(line, "!") {
// 			// files and directories at top-level
// 			if strings.HasPrefix(line, "/") {
// 				// if it starts with "/" it can be a file or dir, then if it
// 				// also ends with "/" then it is a dir
// 				gi.paths = append(gi.paths, dir+line)
// 				continue

// 				// dir or file paths at root-level
// 			} else if strings.Contains(line, "/") {
// 				// like: cmd/main.go - cmd/main - cmd/main/
// 				gi.paths = append(gi.paths, dir+"/"+line)
// 				continue

// 			}
// 		}

// 	}
// }

// func wildcardValidate(path string) []string {
// 	// Supports "*", "**", "???..."
// 	//TODO a combination of ? and * and **
// 	if !strings.Contains(path, "*") {
// 		return nil
// 	}

// 	we := wildcardEscaped(path)
// 	for ch, indxs := range we {

// 	}

// 	if strings.Count(path, "*") == 1 {

// 	} else if strings.Count(path, "*") == 2 {

// 	} else if strings.Contains(path, "?") {

// 	} else {
// 		log.Println("could not handle path %v\n", path)
// 		return nil
// 	}

// }

// func wildcardEscaped(path string) map[string][]int {
// 	return allIndexSubs(path, []string{"*", "**", "?", "!"}, true)
// }

// func allIndexSubs(s string, substrs []string, skipEscaped bool) map[string][]int {
// 	result := make(map[string][]int)

// 	for _, substr := range substrs {
// 		inds := allIndex(s, substr, skipEscaped)
// 		result[substr] = inds
// 	}
// 	return result
// }

// func allIndex(s, substr string, skipEscaped bool) []int {
// 	var indexes []int
// 	offset := 0

// 	for {
// 		i := strings.Index(s[offset:], substr)
// 		if i == -1 {
// 			break
// 		}
// 		finalIndex := offset + i

// 		if s[finalIndex+1:finalIndex+1+utf8.RuneCountInString(substr)] == substr {
// 			continue
// 		}

// 		if s[finalIndex-1] == '\\' {
// 			continue
// 		}

// 		indexes = append(indexes, finalIndex)
// 		offset += i + len(substr)
// 	}
// 	return indexes
// }

// func allIndexMulti(s string, substrs []string, skipEscaped bool) map[string][]int {
// 	result := make(map[string][]int)
// 	for _, substr := range substrs {
// 		result[substr] = allIndex(s, substr, skipEscaped)
// 	}
// 	return result
// }

