package aeryavenue

import "os"

type InputSelector interface {
	SelectItem(keys []string) (string, error)
}

func GetInputSelector() InputSelector {
	// don't prompt for input while in automated pipeline
	envVars := []string{"GITHUB_ACTIONS", "GITLAB_CI"}

	for _, envVar := range envVars {
		s := os.Getenv(envVar)
		b, err := stringToBool(s)
		if err != nil {
			return &TviewInputSelector{}
		}

		if b {
			return &RandomItemInputSelector{}
		}
	}

	return &TviewInputSelector{}
}

func SelectItem(m map[string]string, selector InputSelector) (string, error) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	item, err := selector.SelectItem(keys)
	if err != nil {
		return "", err
	}

	return m[item], nil
}
