package aeryavenue

import (
	"log/slog"

	"github.com/atotto/clipboard"
)

type (
	BlackholeDestination struct{}
	ClipboardDestination struct{}
	ConsoleDestination   struct{}
)

// fixme: never see output of this, dunno why
func (cd *BlackholeDestination) Write(data string) error {
	return nil
}

type OutputDestination interface {
	Write(data string) error
}

type FileDestination struct {
	FilePath string
}

func writeToClipboard(s string) error {
	err := clipboard.WriteAll(s)
	if err != nil {
		slog.Error("error writing to clipboard", "error", err)
	}

	return err
}
