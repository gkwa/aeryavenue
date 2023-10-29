package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/taylormonacelli/aeryavenue"
	"github.com/taylormonacelli/goldbug"
)

var (
	verbose   bool
	logFormat string
)

func main() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output (shorthand)")

	flag.StringVar(&logFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	if verbose || logFormat != "" {
		if logFormat == "json" {
			goldbug.SetDefaultLoggerJson(slog.LevelDebug)
		} else {
			goldbug.SetDefaultLoggerText(slog.LevelDebug)
		}
	}

	paths := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	code := aeryavenue.Main(paths)
	os.Exit(code)
}
