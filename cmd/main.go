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

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "mdme: %v\n", err)
		os.Exit(1)
	}
}

// TODO generate md file with base root folder name
func run() error {
	var verbose bool
	var suppress bool

	flag.BoolVar(&verbose, "verbose", false, "enable debug logging")
	flag.BoolVar(&verbose, "v", false, "alias for -verbose")
	flag.BoolVar(&suppress, "suppress", false, "suppress console output, only copy to clipboard")
	flag.BoolVar(&suppress, "s", false, "alias for -suppress")

	// Current directory and no path provided, immediately a flag
	path := "."
	if len(os.Args) >= 2 && !strings.HasPrefix(os.Args[1], "-") {
		// Provided directory no flag, then other flags
		path = os.Args[1]
		if isDir, err := mdme.IsDir(path); err != nil {
			return fmt.Errorf("Error parsing path: %v", mdme.ErrorMsg(err))
		} else if !isDir {
			// Handle file and return md repr
			return fmt.Errorf("%v: Not a directory", path)
		}
		flag.CommandLine.Parse(os.Args[2:])
	} else {
		flag.Parse()
	}

	logLevel := slog.LevelInfo
	if verbose {
		logLevel = slog.LevelDebug
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	})))

	if path == "." {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Error getting current directory: %v", mdme.ErrorMsg(err))
		}
		path = cwd
	}

	files, err := mdme.ListFiles(path)
	if err != nil {
		return fmt.Errorf("Error parsing files: %v", mdme.ErrorMsg(err))
	}

	md, err := mdme.ToMD(files, path)
	if err != nil {
		return fmt.Errorf("Error parsing files: %v", mdme.ErrorMsg(err))
	}

	// Raise error only if output is supressed, else ignore
	if err := glippy.Set(md); err != nil {
		if suppress {
			return fmt.Errorf("Error pasting output to clipboard: %v", mdme.ErrorMsg(err))
		}
		slog.Warn("clipboard copy failed", "error", mdme.ErrorMsg(err))
	}

	if !suppress {
		fmt.Print(md)
	}

	return nil

}
