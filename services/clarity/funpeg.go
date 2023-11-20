package clarity

import (
	fileModel "c-vod/models/fileModel"
	"c-vod/utils/globals"
	"c-vod/utils/helper"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

/*
Funpeg will handle the video format standardization
*/
type Funpeg struct{}

func (fu *Funpeg) standardization(file *fileModel.File) error {

	// + funpeg will convert video into 480p and 720p res files
	// in stage2 folder
	//err := fu.mockStandardization(file)
	err := fu._standardization(file)

	if err != nil {
		// TODO: rollback
		return err
	}

	// + funpeg will call trauma to delete step1 file
	err = Trauma.deleteUploadedFile(file)

	if err != nil {
		// TODO: rollback
		return err
	}

	// + funpeg will call edborn to update dbfile status to step2
	err = Edborn.updateFileStatus(file, fileModel.NORMALIZED)

	if err != nil {
		// TODO: rollback
		return err
	}

	return nil
}

func (fu *Funpeg) _standardization(file *fileModel.File) error {

	var ffmpeg string

	switch runtime.GOOS {
	case "darwin":
		ffmpeg = "./ffmpeg-osx"
	case "win32":
		ffmpeg = "./ffmpeg.exe"
	case "linux":
		ffmpeg = "./ffmpeg-linux"
	}

	file_id := fmt.Sprintf("%d", file.Id)

	video_input := "./../storage/stage1/" + file_id + "." + file.Ext
	scale480 := "scale=720:480"
	scale720 := "scale=1280:720"
	video480_output := "./../storage/stage2/" + file_id + "_480." + file.Ext
	video720_output := "./../storage/stage2/" + file_id + "_720." + file.Ext

	cmd480 := exec.Command(ffmpeg,
		"-i",
		video_input,
		"-vf",
		scale480,
		"-c:v",
		"libx264",
		"-crf",
		"23",
		"-preset",
		"medium",
		"-c:a",
		"aac",
		"-b:a",
		"128k",
		video480_output,
	)

	cmd480.Dir = filepath.Join(helper.GetCurrentDir(), "ffmpeg")

	if globals.App.Config.Log_enabled == "true" {
		// showing ffmpeg output in terminal
		cmd480.Stdout = os.Stdout
		cmd480.Stderr = os.Stderr
	}

	var err error = nil

	// + elisma will store output in a new folder with dbfile id in stage3 folder
	err = cmd480.Run()

	if err != nil {
		//TODO: rollback
		fmt.Println("packager cmd error :", err.Error())
		return err
	}

	cmd720 := exec.Command(ffmpeg,
		"-i",
		video_input,
		"-vf",
		scale720,
		"-c:v",
		"libx264",
		"-crf",
		"23",
		"-preset",
		"medium",
		"-c:a",
		"aac",
		"-b:a",
		"128k",
		video720_output,
	)

	cmd720.Dir = filepath.Join(helper.GetCurrentDir(), "ffmpeg")

	if globals.App.Config.Log_enabled == "true" {
		// showing ffmpeg output in terminal
		cmd720.Stdout = os.Stdout
		cmd720.Stderr = os.Stderr
	}

	// + elisma will store output in a new folder with dbfile id in stage3 folder
	err = cmd720.Run()

	if err != nil {
		//TODO: rollback
		fmt.Println("packager cmd error :", err.Error())
		return err
	}

	return nil
}

func (fu *Funpeg) mockStandardization(file *fileModel.File) error {

	file_id := fmt.Sprintf("%d", file.Id)

	source_path := filepath.Join(helper.GetCurrentDir(), "./storage/stage1/"+file_id+"."+file.Ext)

	sourceFile, err := os.Open(source_path)

	if err != nil {
		//TODO: rollback
		log.Printf("could not open file : %v", err)
		return err
	}

	defer sourceFile.Close()

	dest_720 := filepath.Join(helper.GetCurrentDir(), "./storage/stage2/"+file_id+"_720."+file.Ext)

	dest_720File, err := os.Create(dest_720)

	if err != nil {
		//TODO: rollback
		log.Printf("could not create file : %v", err)
		return err
	}
	defer dest_720File.Close()

	_, err = io.Copy(dest_720File, sourceFile)

	if err != nil {
		//TODO: rollback
		log.Printf("could not copy data : %v", err)
		return err
	}

	dest_480 := filepath.Join(helper.GetCurrentDir(), "./storage/stage2/"+file_id+"_480."+file.Ext)

	dest_480File, err := os.Create(dest_480)

	if err != nil {
		//TODO: rollback
		log.Printf("could not create file : %v", err)
		return err
	}

	defer dest_480File.Close()

	// this seeks back to the start point of sourse file
	sourceFile.Seek(0, 0)

	_, err = io.Copy(dest_480File, sourceFile)

	if err != nil {
		//TODO: rollback
		log.Printf("could not copy data : %v", err)
		return err
	}

	log.Print("copy was compeleted!")

	return nil
}

////./ffmpeg-mac -i input.mp4 -vf "scale=1280:720" -c:v libx264 -crf 23 -preset medium -c:a aac -b:a 128k output.mp4
