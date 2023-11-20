package services

import (
	fileModel "c-vod/models/fileModel"
	uploadModel "c-vod/models/uploadModel"
	"c-vod/utils/globals"
	"c-vod/utils/helper"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func FindUploadRecord(upload_token string, upload_id string) (*uploadModel.Upload, error) {

	upload := &uploadModel.Upload{}

	int_id, err := strconv.Atoi(upload_id)

	if err != nil {
		return nil, errors.New("invalid upload_id")
	}

	result := globals.App.DB.Where(
		"id = ? AND token = ?",
		int_id,
		upload_token,
	).Find(upload)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("upload record not found")
	}

	return upload, nil
}

func UploadExpirationCheck(upload *uploadModel.Upload) (bool, error) {

	if upload.Status == uploadModel.Expired {
		return true, nil
	}

	if upload.ExpiresAt < time.Now().UnixMilli() {

		uq_result := globals.App.DB.Model(&uploadModel.Upload{}).
			Where("id = ?", upload.Id).
			Update("status", uploadModel.Expired)

		if uq_result.Error != nil {
			return false, uq_result.Error
		}

		return true, nil
	}

	return false, nil
}

func StoreFile(file_h *multipart.FileHeader, upload *uploadModel.Upload) error {

	// check if fileHeader can open without error
	file, err := file_h.Open()
	if err != nil {
		return err
	}

	// compare file size
	if upload.Size != file_h.Size {
		return errors.New("file size does not match")
	}

	//compate file ext
	file_ext, err := helper.GetFileExtension(file_h)
	if err != nil {
		return err
	}
	if upload.Ext != file_ext {
		return errors.New("file size extension not match")
	}

	//TODO: check if file is truly the same type

	//set file status base on file type
	file_status := fileModel.UPLOADED
	if upload.Type == fileModel.OTHER {
		file_status = fileModel.READY_FOR_STORAGE
	}

	//create the file record
	db_file := fileModel.File{
		Id:       upload.Id,
		Type:     upload.Type,
		Ext:      upload.Ext,
		Status:   file_status,
		Size:     upload.Size,
		Duration: 0,
	}
	result := globals.App.DB.Create(&db_file)
	if result.Error != nil {
		//TODO: rollback
		return result.Error
	}

	file_id := fmt.Sprint(db_file.Id)

	//store the file in temp storage dir
	destination_path := ""
	if db_file.Type == fileModel.VIDEO {
		destination_path = "./storage/stage1/" + file_id + "." + db_file.Ext
	} else {

		dir_path := filepath.Join(helper.GetCurrentDir(), "./storage/stage3/"+file_id)
		err = helper.CreateDirectory(dir_path)
		if err != nil {
			//TODO: rollback
			return err
		}
		destination_path = "./storage/stage3/" + file_id + "/" + file_id + "." + db_file.Ext
	}
	destination_path = filepath.Join(helper.GetCurrentDir(), destination_path)
	if err != nil {
		//TODO: rollback -> delete file record from db
		return err
	}
	destination_file, err := os.Create(destination_path)
	if err != nil {
		//TODO: rollback -> delete file record from db
		return err
	}
	if _, err = io.Copy(destination_file, file); err != nil {
		//TODO: rollback -> delete file record from db
		return err
	}

	return nil
}
