package aeryavenue

import (
	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TviewInputSelector struct{}

func (selector *TviewInputSelector) SelectItem(listItems []string) (string, error) {
	if len(listItems) < 1 {
		return "", nil
	}

	app := tview.NewApplication()

	var selectedItem string

	// Create a list widget and add the items to it
	list := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
			selectedItem = listItems[index]
			returnValue(selectedItem, &BlackholeDestination{})
			// returnValue(selectedItem, &ClipboardDestination{})
			// returnValue(selectedItem, &ConsoleDestination{})
			// returnValue(selectedItem, &FileDestination{FilePath: "items.txt"})
			app.Stop()
		})

	for _, item := range listItems {
		list.AddItem(item, "", rune(0), nil)
	}

	// Set up key bindings to navigate the list
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'n':
				list.SetCurrentItem((list.GetCurrentItem() + 1) % list.GetItemCount())
				return nil
			case 'p':
				current := list.GetCurrentItem()
				if current == 0 {
					current = list.GetItemCount()
				}
				list.SetCurrentItem((current - 1) % list.GetItemCount())
				return nil
			case 'q':
				app.Stop()
				return nil
			}
		}

		return event
	})

	// Set the list widget as the root and run the application
	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	return selectedItem, nil
}

func returnValue(val string, output OutputDestination) {
	slog.Debug("out", "val", val)

	if err := output.Write(val); err != nil {
		slog.Error("error writing to output", "error", err)
	}
}
