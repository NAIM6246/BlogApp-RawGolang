package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (h *PostHandlers) uploadFile(r *http.Request) (string, error) {
	err := r.ParseMultipartForm(10 << 20) // Parse multipart form data with 10MB max memory
	if err != nil {
		return "", err
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a new file in the uploads directory
	f, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Copy the file contents to the new file
	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("http://localhost:8005/files/%s", handler.Filename)
	return fileURL, nil
}
