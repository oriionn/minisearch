package utils

import (
	"os"
	"slices"
)

func DevMode() bool {
	return slices.Contains(os.Args, "--dev")
}
