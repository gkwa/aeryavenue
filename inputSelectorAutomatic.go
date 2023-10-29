package aeryavenue

type AutomaticItemInputSelector struct{}

func (selector *AutomaticItemInputSelector) SelectItem(listItems []string) (string, error) {
	return listItems[0], nil
}
