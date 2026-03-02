# mdme (markdown me)

A simple cli tool to generate code and text blocks for viewing, sharing with people or with an LLM.

## Installation

```bash
go install github.com/alafeefidev/mdme/cmd/mdme@latest
```

## Usage

```bash
mdme <path> [flags]
```

## Examples
1. Get files from current directory with console output suppressed
```bash
mdme -s 
```
2. Get files from provided directory with verbosity of logs set to debug
```bash
mdme /examples -v
```

## Flags
- `-s, --suppress` suppress console output, only copy to clipboard
- `-v, --verbose` enable debug logging
- `-d, --depth` max directories depth (0 = unlimited)
- `--max-files` max number of files to process

## Features
- Support for converting files in chosen directory to md code representations
- Ignores files in .gitignore. With support for .mdignore for files to ignore only from mdme

## Planned Features
- Flags to overwrite default ignore list, ignore .gitignore, adding regex matching for ignoring files.
- To work on files too, now it only accepts directories
- Print to an .md file
- and more...