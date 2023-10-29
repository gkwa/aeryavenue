package aeryavenue

import "os"

type InputSelector interface {
	SelectItem(keys []string) (string, error)
}

func GetInputSelector() InputSelector {
	randSelector := &RandomItemInputSelector{}
	tviewSelector := &TviewInputSelector{}

	// don't prompt for input while in automated pipeline
	envVars := []string{"GITHUB_ACTIONS", "GITLAB_CI"}

	for _, envVar := range envVars {
		s := os.Getenv(envVar)
		b, err := stringToBool(s)
		if err != nil {
			return tviewSelector
		}

		if b {
			return randSelector
		}
	}

	return tviewSelector
}

func SelectItem(items map[string]string, selector InputSelector) (string, error) {
	sortedKeys := sortedKeys(items)
	item, err := selector.SelectItem(sortedKeys)
	if err != nil {
		return "", err
	}

	return items[item], nil
}
