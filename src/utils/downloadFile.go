package utils

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(params DownloadFileParams) error {
	if !strings.HasSuffix(params.DownloadPath, "/") {
		params.DownloadPath = params.DownloadPath + "/"
	}

	resp, err := http.Get(params.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err := os.Stat(params.DownloadPath); os.IsNotExist(err) {
		return os.MkdirAll(params.DownloadPath, 0755)
	}

	out, err := os.Create(params.DownloadPath + params.Filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

type DownloadFileParams struct {
	Filename     string
	URL          string
	DownloadPath string
}
