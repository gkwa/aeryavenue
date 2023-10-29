package aeryavenue

type AutomaticItemInputSelector struct{}

func (selector *AutomaticItemInputSelector) SelectItem(listItems []string) (string, error) {
	if len(listItems) < 1 {
		return "", nil
	}

	return listItems[0], nil
}
