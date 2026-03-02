package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/alafeefidev/mdme"
	"github.com/f1bonacc1/glippy"
)

// TODO add copy to keyboard
func main() {
	var verbose bool
	var suppress bool

	flag.BoolVar(&verbose, "verbose", false, "enable debug logging")
	flag.BoolVar(&verbose, "v", false, "enable debug logging (shorthand)")
	flag.BoolVar(&suppress, "suppress", false, "suppress console output, only copy to clipboard")
	flag.BoolVar(&suppress, "s", false, "suppress console output, only copy to clipboard (shorthand)")

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

	logLevel := slog.LevelInfo
	if verbose {
		logLevel = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

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

	// Raise error only if output is supressed, else ignore
	err = glippy.Set(md)
	if !suppress {
		fmt.Print(md)
	} else if suppress && err != nil {
		fmt.Fprintf(os.Stderr, "Error pasting output to clipboard: %v\n", mdme.ErrorMsg(err))
		os.Exit(1)
	}
	
	
}
