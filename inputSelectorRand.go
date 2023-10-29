package aeryavenue

import (
	"math/rand"
	"time"
)

type (
	RandomItemInputSelector struct{}
)

func (selector *RandomItemInputSelector) SelectItem(keys []string) (string, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rng.Intn(len(keys))

	return keys[index], nil
}
