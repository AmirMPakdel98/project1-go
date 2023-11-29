package clarity

import (
	"c-vod/models/enckeyModel"
	fileModel "c-vod/models/fileModel"
	"c-vod/utils/globals"
	"c-vod/utils/types"
	"fmt"

	"gorm.io/gorm"
)

/*
Edborn will handle the database jobs
*/
type _Edborn struct{}

var Edborn *_Edborn

func (ed *_Edborn) findWaitingFile() *fileModel.File {
	dbfile := &fileModel.File{}

	result := globals.App.DB.Where(
		"status IN (?,?,?)",
		fileModel.UPLOADED,
		fileModel.NORMALIZED,
		fileModel.READY_TO_STORE,
	).First(dbfile)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		fmt.Println("Error in finding File in db:", result.Error)

		return nil
	}

	if result.RowsAffected == 0 {
		return nil
	}

	return dbfile
}

func (ed *_Edborn) updateFileStatus(file *fileModel.File, status fileModel.Status) error {

	uq_result := globals.App.DB.Model(&fileModel.File{}).Where("id = ?", file.Id).Update("status", status)

	if uq_result.Error != nil {
		fmt.Println("error on updating file's status in db:", uq_result.Error.Error())
		return uq_result.Error
	}

	return nil
}

func (ed *_Edborn) updateFileDuration(file *fileModel.File, duration int) error {

	uq_result := globals.App.DB.Model(&fileModel.File{}).
		Where("id = ?", file.Id).
		Update("duration", duration)

	if uq_result.Error != nil {
		fmt.Println("error on updating file's duration in db:", uq_result.Error.Error())
		return uq_result.Error
	}

	return nil
}

func (ed *_Edborn) insertVideoKeys(keys *types.VideoKeys, file_id uint32) error {

	newKeys := &enckeyModel.Enckey{
		File_id: file_id,
		Payload: keys.String(),
	}

	result := globals.App.DB.Create(newKeys)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
