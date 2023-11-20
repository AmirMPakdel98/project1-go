package helper

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

func GetFileExtension(fileHeader *multipart.FileHeader) (string, error) {

	filename := fileHeader.Filename

	extension := path.Ext(filename)[1:]

	if extension == "" {
		return "", errors.New("unable to determine file extension")
	}

	return extension, nil
}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

func MoveDirectory(sourcePath, destinationPath string) error {
	// Check if the source directory exists
	sourceInfo, err := os.Stat(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to get source directory info: %v", err)
	}
	if !sourceInfo.IsDir() {
		return fmt.Errorf("source path is not a directory")
	}

	// Create the destination directory if it doesn't exist
	err = os.MkdirAll(destinationPath, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to create destination directory: %v", err)
	}

	// Walk through the source directory and move each file to the destination directory
	err = filepath.Walk(sourcePath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access file or directory: %v", err)
		}

		// Skip the source directory itself
		if filePath == sourcePath {
			return nil
		}

		// Construct the destination file path
		destinationFilePath := filepath.Join(destinationPath, filePath[len(sourcePath):])

		// If the current item is a directory, create it in the destination directory
		if fileInfo.IsDir() {
			err = os.MkdirAll(destinationFilePath, fileInfo.Mode())
			if err != nil {
				return fmt.Errorf("failed to create destination directory: %v", err)
			}
		} else {
			// If the current item is a file, move it to the destination directory
			err = MoveFile(filePath, destinationFilePath)
			if err != nil {
				return fmt.Errorf("failed to move file: %v", err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to move directory: %v", err)
	}

	return nil
}

func MoveFile(sourcePath, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to remove source file: %v", err)
	}

	return nil
}

func CreateDirectory(dir_path string) error {

	if err := os.Mkdir(dir_path, os.ModePerm); err != nil {
		return err
	}

	return nil
}
