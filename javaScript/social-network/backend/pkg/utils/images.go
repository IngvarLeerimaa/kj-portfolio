package utils

import (
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gofrs/uuid"
)

func SaveImage(file multipart.File, header *multipart.FileHeader, err error) string {
	if err != nil {
		return ""
	}
	defer file.Close()

	fileType := strings.Split(header.Header.Get("Content-Type"), "/")[1]
	fileName := uuid.Must(uuid.NewV4()).String() + "." + fileType
	filePath := "./data/images/" + fileName

	outFile, err := os.Create(filePath)
	if err != nil {
		return ""
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return ""
	}

	return fileName
}
