package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/alafeefidev/mdme"
)

func main() {

	// Current directory and no path provided, immediately a flag
	path := "."
	if len(os.Args) >= 2 {
		// Provided directory no flag, then other flags
		if !strings.HasPrefix(os.Args[1], "-") {
			path = os.Args[1]
			if isDir, err := mdme.IsDir(path); err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing path: %v\n", mdme.ErrorMsg(err))
				os.Exit(1)
			} else if !isDir {
				// Handle file and return md repr
				fmt.Fprintf(os.Stderr, "%v: Not a directory\n", path)
				os.Exit(1)
			}

			flag.CommandLine.Parse(os.Args[2:])
		} else {
			flag.Parse()
		}
	}

	files, err := mdme.ListFiles(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing files: %v\n", mdme.ErrorMsg(err))
		os.Exit(1)
	}

	md, err := mdme.ToMD(files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing files: %v\n", mdme.ErrorMsg(err))
		os.Exit(1)
	}

	fmt.Print(md)
}
