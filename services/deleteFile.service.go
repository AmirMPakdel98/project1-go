package services

import (
	fileModel "c-vod/models/fileModel"
	"c-vod/utils/globals"
	"errors"
	"fmt"
)

func DeleteFile(file *fileModel.File) error {

	file_id := fmt.Sprintf("%d", file.Id)

	object_name := "edu-arch/" + file_id

	err := globals.App.Storage.DeleteDirectory(object_name)

	if err != nil {
		//TODO: rollback
		return err
	}

	//delete file db record
	status := fileModel.DELETED
	del_result := globals.App.DB.Model(&fileModel.File{}).
		Where("id = ?", file.Id).
		Update("status", status)

	if del_result.Error != nil {
		//TODO: rollback
		return del_result.Error
	}

	if del_result.RowsAffected == 0 {
		//TODO: rollback
		return errors.New("error on changing status to delete for file record")
	}

	return nil
}
