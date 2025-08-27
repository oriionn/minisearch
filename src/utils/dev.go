package utils

import (
	"os"
	"slices"
)

func DevMode() bool {
	return slices.Contains(os.Args, "--dev") || slices.Contains(os.Args, "-d")
}
