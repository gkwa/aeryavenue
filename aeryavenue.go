package aeryavenue

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Main(itemsMap map[string]string) int {
	inputSelector := GetInputSelector()
	selectedItem, err := selectItem(itemsMap, inputSelector)
	if err != nil {
		slog.Error("selectItem failed", "error", err)
		return 1
	}

	fmt.Println(selectedItem)

	return 0
}

type (
	BlackholeDestination struct{}
	ClipboardDestination struct{}
	ConsoleDestination   struct{}
)

type (
	RandomItemInputSelector struct{}
	TviewInputSelector      struct{}
)

func selectItem(items map[string]string, is InputSelector) (string, error) {
	sortedKeys := sortedKeys(items)
	item, err := is.SelectItem(sortedKeys)
	if err != nil {
		return "", err
	}

	return items[item], nil
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// fixme: never see output of this, dunno why
func (cd *BlackholeDestination) Write(data string) error {
	return nil
}

type InputSelector interface {
	SelectItem(keys []string) (string, error)
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

func (uis *TviewInputSelector) SelectItem(keys []string) (string, error) {
	app := tview.NewApplication()

	var selectedItem string

	// Create a list widget and add the items to it
	list := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
			selectedItem = keys[index]
			returnValue(selectedItem, &BlackholeDestination{})
			// returnValue(selectedItem, &ClipboardDestination{})
			// returnValue(selectedItem, &ConsoleDestination{})
			// returnValue(selectedItem, &FileDestination{FilePath: "items.txt"})
			app.Stop()
		})
	for _, item := range keys {
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

func stringToBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return b, nil
}

func GetInputSelector() InputSelector {
	ris := &RandomItemInputSelector{}
	uis := &TviewInputSelector{}

	// don't prompt for input while in automated pipeline
	envVars := []string{"GITHUB_ACTIONS", "GITLAB_CI"}

	for _, envVar := range envVars {
		s := os.Getenv(envVar)
		b, err := stringToBool(s)
		if err != nil {
			return uis
		}
		if b {
			return ris
		}
	}

	return uis
}

func (selector *RandomItemInputSelector) SelectItem(keys []string) (string, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rng.Intn(len(keys))
	return keys[index], nil
}
