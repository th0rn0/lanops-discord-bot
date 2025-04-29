package utils

import (
	"strings"
)

func IsImage(filename string) bool {
	lower := strings.ToLower(filename)
	return strings.HasSuffix(lower, ".jpg") || strings.HasSuffix(lower, ".jpeg") ||
		strings.HasSuffix(lower, ".png")
}

func IsMedia(filename string) bool {
	lower := strings.ToLower(filename)

	// Image extensions
	if strings.HasSuffix(lower, ".jpg") || strings.HasSuffix(lower, ".jpeg") ||
		strings.HasSuffix(lower, ".png") {
		return true
	}

	// Video extensions
	if strings.HasSuffix(lower, ".mp4") || strings.HasSuffix(lower, ".mov") ||
		strings.HasSuffix(lower, ".webm") || strings.HasSuffix(lower, ".avi") ||
		strings.HasSuffix(lower, ".mkv") {
		return true
	}

	return false
}
