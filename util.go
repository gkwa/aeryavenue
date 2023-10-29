package aeryavenue

import (
	"strconv"
)

func stringToBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	
	return b, nil
}
