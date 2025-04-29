package utils

import (
	"context"
	"encoding/base64"
	"os"
	"path/filepath"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func UploadToDrive(filePath string, parentFolderID string) (string, error) {
	ctx := context.Background()
	googleCreds := os.Getenv("GOOGLE_KEY_JSON_BASE64")

	jsonBytes, err := base64.StdEncoding.DecodeString(googleCreds)
	if err != nil {
		return "", err
	}

	// Authenticate using the service account
	googleConfig, err := google.JWTConfigFromJSON(jsonBytes, drive.DriveFileScope)
	if err != nil {
		return "", err
	}

	googleClient := googleConfig.Client(ctx)

	// Create Drive service
	srv, err := drive.NewService(ctx, option.WithHTTPClient(googleClient))
	if err != nil {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	f := &drive.File{
		Name:    filepath.Base(filePath),
		Parents: []string{parentFolderID},
	}

	createdFile, err := srv.Files.Create(f).Media(file).Do()
	if err != nil {
		return "", err
	}
	return createdFile.Id, nil
}
