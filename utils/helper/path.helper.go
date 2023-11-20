package helper

import (
	"log"
	"os"
	"path/filepath"
)

func GetCurrentDir() string {

	from := os.Getenv("ENV")

	if from == "dev" {

		path, err := os.Getwd()
		if err != nil {
			log.Println("failed to GetCurrentDir :" + err.Error())
			return "."
		}
		return path

	} else {
		ex, err := os.Executable()

		if err != nil {
			log.Println("failed to GetCurrentDir :" + err.Error())
			return "."
		}
		return filepath.Dir(ex)
	}

}
