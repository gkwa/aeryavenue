package aeryavenue

import (
	"fmt"
	"log/slog"
)

func Main(itemsMap map[string]string) int {
	inputSelector := GetInputSelector()
	selectedItem, err := SelectItem(itemsMap, inputSelector)
	if err != nil {
		slog.Error("selectItem failed", "error", err)
		return 1
	}

	fmt.Println(selectedItem)

	return 0
}
