package aeryavenue

import "log/slog"

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

func returnValue(val string, output OutputDestination) {
	slog.Debug("out", "val", val)

	if err := output.Write(val); err != nil {
		slog.Error("error writing to output", "error", err)
	}
}
