package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type LocalClient struct {
	BaseURL   string
	UploadDir string
}

func NewLocalClient(baseURL, uploadDir string) Provider {
	return &LocalClient{
		BaseURL:   baseURL,
		UploadDir: uploadDir,
	}
}

func (l *LocalClient) UploadFile(ctx context.Context, filename string, data []byte) (string, error) {
	// Ensure upload directory exists
	// Filename might contain subdirectories (e.g. images/...) so we need to ensure the dir structure exists
	fullPath := filepath.Join(l.UploadDir, filename)
	dir := filepath.Dir(fullPath)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		logrus.Errorf("failed to create upload directory: %v", err)
		return "", err
	}

	// Create/Overwrite file
	file, err := os.Create(fullPath)
	if err != nil {
		logrus.Errorf("failed to create file: %v", err)
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		logrus.Errorf("failed to write data: %v", err)
		return "", err
	}

	// For local, we assume the file is served via the BaseURL + filename
	// e.g. http://localhost:8001/public/images/filename.jpg
	// We ensure consistency with how GCS returns URLs
	url := fmt.Sprintf("%s/%s", l.BaseURL, filename)
	return url, nil
}
