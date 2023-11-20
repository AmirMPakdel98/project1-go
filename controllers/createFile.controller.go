package controllers

import (
	"c-vod/models/fileModel"
	"c-vod/services"
	"c-vod/utils/response"
	"c-vod/utils/types"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateFile(c *fiber.Ctx) error {

	body := &types.CreateFilesReq{}

	err := c.BodyParser(body)

	if err != nil {
		response.Error(c, -1, err)
		return nil
	}

	// check if file is not biger than 900MB
	if body.File_size > 943718400 {
		err = errors.New("file size cannot be more than 943718400 bytes")
		response.Error(c, -2, err)
		return nil
	}

	// check if file type VIDEO has mp4 ext
	if body.File_type == int(fileModel.VIDEO) && body.File_ext != "mp4" {
		err = errors.New("file type VIDEO must have mp4 file ext")
		response.Error(c, -3, err)
		return nil
	}

	// check if file type OTHER doesn't have mp4 ext
	if body.File_type == int(fileModel.OTHER) && body.File_ext == "mp4" {
		err = errors.New("file type OTHER cannot have mp4 file ext")
		response.Error(c, -3, err)
		return nil
	}

	upload, err := services.CreateUpload(body)

	if err != nil {
		response.ServerError(c, err)
		return nil
	}

	upload_url := services.GetUploadURL(upload)

	response.Success(
		c,
		&types.CreateFilesResData{
			Object_id:  upload.Id,
			Upload_url: upload_url,
		},
	)

	return nil
}
