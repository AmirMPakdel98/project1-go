package services

import (
	fileModel "c-vod/models/fileModel"
	"c-vod/utils/globals"
	"c-vod/utils/types"
	"errors"
	"fmt"
)

func FindFileRecord(file_id int) (*fileModel.File, error) {

	file := &fileModel.File{}

	result := globals.App.DB.Where("id = ?", file_id).First(file)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		//TODO: rollback
		return nil, errors.New("file not found")
	}

	return file, nil
}

func GetFileResponseData(file *fileModel.File) (*types.GetFilesResData, error) {

	file_id := fmt.Sprintf("%d", file.Id)

	file_name := ""

	var keys []types.FileKey

	if (file.Type == fileModel.VIDEO) && (file.Status == fileModel.COMPLETED) {
		file_name = fmt.Sprintf(
			"__0000_ZGFzaF9zdHJlYW1fc2VnbWVudF9maWxl_%s"+
				"_dash_stream_manifest.mpd",
			file_id,
		)
		keys = []types.FileKey{
			{Id: "f3c5e0361e6654b28f8049c778b23946", Value: "a4631a153a443df9eed0593043db7519"},
			{Id: "abba271e8bcf552bbd2e86a434a9a5d9", Value: "69eaa802a6763af979e8d1940fb88392"},
			{Id: "6d76f25cb17f5e16b8eaef6bbf582d8e", Value: "cb541084c99731aef4fff74500c12ead"},
		}
	} else if (file.Type == fileModel.OTHER) && (file.Status == fileModel.COMPLETED) {
		file_name = fmt.Sprintf("%s.%s", file_id, file.Ext)
	}

	url := fmt.Sprintf(
		"https://vod-storage-test.s3.ir-thr-at1.arvanstorage.ir"+
			"/edu-arch/%s/%s",
		file_id,
		file_name,
	)

	return &types.GetFilesResData{
		File_status:   int(file.Status),
		File_duration: file.Duration,
		File_url:      url,
		File_keys:     keys,
	}, nil
}
