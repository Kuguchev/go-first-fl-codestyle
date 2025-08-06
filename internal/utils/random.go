package utils

import (
	"math/rand"
)

func Random(min, max int) int {
	if min > max {
		min, max = max, min
	}

	return rand.Intn(max-min) + min
}
