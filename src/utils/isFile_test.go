package utils

import "testing"

func TestIsImage(t *testing.T) {
	cases := []struct {
		filename string
		want     bool
	}{
		{"photo.jpg", true},
		{"photo.jpeg", true},
		{"photo.png", true},
		{"PHOTO.JPG", true},
		{"PHOTO.PNG", true},
		{"video.mp4", false},
		{"document.pdf", false},
		{"archive.zip", false},
		{"", false},
	}
	for _, tc := range cases {
		got := IsImage(tc.filename)
		if got != tc.want {
			t.Errorf("IsImage(%q) = %v, want %v", tc.filename, got, tc.want)
		}
	}
}

func TestIsMedia(t *testing.T) {
	cases := []struct {
		filename string
		want     bool
	}{
		// Images
		{"photo.jpg", true},
		{"photo.jpeg", true},
		{"photo.png", true},
		// Videos
		{"clip.mp4", true},
		{"clip.mov", true},
		{"clip.webm", true},
		{"clip.avi", true},
		{"clip.mkv", true},
		// Case insensitive
		{"CLIP.MP4", true},
		{"PHOTO.PNG", true},
		// Non-media
		{"document.pdf", false},
		{"archive.zip", false},
		{"script.sh", false},
		{"", false},
	}
	for _, tc := range cases {
		got := IsMedia(tc.filename)
		if got != tc.want {
			t.Errorf("IsMedia(%q) = %v, want %v", tc.filename, got, tc.want)
		}
	}
}
