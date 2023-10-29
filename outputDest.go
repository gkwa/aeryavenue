package aeryavenue

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
