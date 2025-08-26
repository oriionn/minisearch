package utils

import (
	"math/rand/v2"
)

func Random(min int, max int) int {
	return rand.IntN(max-min+1) + min
}
