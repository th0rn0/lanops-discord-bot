package utils

import (
	"strings"
)

func IsImage(filename string) bool {
	lower := strings.ToLower(filename)
	return strings.HasSuffix(lower, ".jpg") ||
		strings.HasSuffix(lower, ".jpeg") ||
		strings.HasSuffix(lower, ".png")
}
