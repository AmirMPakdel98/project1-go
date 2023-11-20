package clarity

/*
	trauma will handle the file transition
*/

import (
	fileModel "c-vod/models/fileModel"
	"c-vod/utils/helper"
	"fmt"
	"os"
	"path/filepath"
)

type _Trauma struct{}

var Trauma *_Trauma

func (tr *_Trauma) deleteUploadedFile(file *fileModel.File) error {

	file_id := fmt.Sprintf("%d", file.Id)

	rm_err := os.Remove(filepath.Join(
		helper.GetCurrentDir(),
		"./storage/stage1/"+file_id+"."+file.Ext))

	if rm_err != nil {
		fmt.Println("error on removing file from stage1 :", rm_err.Error())
		return rm_err
	}

	return nil
}

func (tr *_Trauma) deleteNormalizedVideoSource(file *fileModel.File) error {

	file_id := fmt.Sprintf("%d", file.Id)

	rm_err1 := os.Remove(filepath.Join(
		helper.GetCurrentDir(),
		"./storage/stage2/"+file_id+"_480."+file.Ext))

	rm_err2 := os.Remove(filepath.Join(
		helper.GetCurrentDir(),
		"./storage/stage2/"+file_id+"_720."+file.Ext))

	if rm_err1 != nil {
		return rm_err1
	}

	if rm_err2 != nil {
		return rm_err2
	}

	return nil
}

func (tr *_Trauma) deleteObjectDirectory(file *fileModel.File) error {

	file_id := fmt.Sprintf("%d", file.Id)

	rm_err := os.RemoveAll(
		filepath.Join(helper.GetCurrentDir(), "./storage/stage3/"+file_id),
	)

	if rm_err != nil {
		return rm_err
	}

	return nil
}
