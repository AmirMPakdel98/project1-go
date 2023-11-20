package services

import (
	fileModel "c-vod/models/fileModel"
	uploadModel "c-vod/models/uploadModel"
	"c-vod/utils/globals"
	"c-vod/utils/types"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CreateUpload(body *types.CreateFilesReq) (*uploadModel.Upload, error) {

	token := uuid.New().String()
	ext := body.File_ext
	size := body.File_size
	expiresAt := time.Now().Add(time.Hour).UnixMilli()

	upload := &uploadModel.Upload{
		Token:     token,
		Type:      fileModel.Type(body.File_type),
		Ext:       ext,
		Size:      int64(size),
		ExpiresAt: expiresAt,
	}

	result := globals.App.DB.Create(upload)

	if result.Error != nil {
		return nil, result.Error
	}

	return upload, nil
}

func GetUploadURL(upload *uploadModel.Upload) string {

	//TODO: use dynamic link
	domain := globals.App.Config.App_domain

	prefix := globals.App.Config.App_api_prefix_v1

	expiration := upload.ExpiresAt

	return fmt.Sprintf(domain+prefix+"/file/upload?t=%s&i=%d&expire_timestamp=%d",
		upload.Token, int(upload.Id), expiration)
}
