package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
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

// Skip by extension like pdf files
// TODO generate md file with base root folder name
func run() error {
	var verbose bool
	var suppress bool
	var maxDepth int
	var maxFiles int
	var output string

	flag.BoolVar(&verbose, "verbose", false, "enable debug logging")
	flag.BoolVar(&verbose, "v", false, "alias for -verbose")
	flag.BoolVar(&suppress, "suppress", false, "suppress console output, only copy to clipboard")
	flag.BoolVar(&suppress, "s", false, "alias for -suppress")
	flag.IntVar(&maxDepth, "depth", 50, "max directories depth (0 = unlimited)")
	flag.IntVar(&maxDepth, "d", 50, "alias for -depth")
	flag.IntVar(&maxFiles, "max-files", 100, "max number of files to process")
	flag.StringVar(&output, "output", "", "print md file to this file name")
	flag.StringVar(&output, "o", "", "alias for -output")

	// Current directory and no path provided, immediately a flag
	path := "."
	if len(os.Args) >= 2 && !strings.HasPrefix(os.Args[1], "-") {
		// Provided directory no flag, then other flags
		path = os.Args[1]

		isDir, err := mdme.IsDir(path)
		if err != nil {
			return fmt.Errorf("Error parsing path: %v", mdme.ErrorMsg(err))
		}

		if !isDir {
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

	abs, _ := filepath.Abs(path)
	if mdme.IsHomeDir(abs) {
		fmt.Fprintf(os.Stderr,
			"warning: scanning home directory %s, this will take a while\n it is recommended to use -d to limit depth or choose a directory\n", abs)
	}

	files, err := mdme.ListFiles(path, maxDepth, maxFiles)
	if err != nil {
		return fmt.Errorf("Error parsing files: %v", mdme.ErrorMsg(err))
	}

	md, err := mdme.ToMD(files, path)
	if err != nil {
		return fmt.Errorf("Error converting files to MD: %v", mdme.ErrorMsg(err))
	}

	// Raise error only if output is supressed, else ignore
	md = strings.ReplaceAll(md, "\x00", "") // To fix string with NUL passed error for UTF16PtrFromString, hmmm

	if err := glippy.Set(md); err != nil {
		if suppress {
			return fmt.Errorf("Error pasting output to clipboard: %v", mdme.ErrorMsg(err))
		}
		slog.Warn("clipboard copy failed", "error", mdme.ErrorMsg(err))
	}

	if !suppress {
		fmt.Print(md)
	}

	if output != "" {
		if err := mdme.CreateMDFile(output, path, md); err != nil {
			return fmt.Errorf("Error writing md file (%s): %v", output, mdme.ErrorMsg(err))
		}
	}

	return nil

}
