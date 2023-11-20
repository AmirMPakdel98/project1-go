package controllers

import (
	"c-vod/models/uploadModel"
	"c-vod/services"
	"c-vod/utils/helper"
	"c-vod/utils/response"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {

	token := c.Query("t", "")
	upload_id := c.Query("i", "")

	var err error

	if token == "" || upload_id == "" {
		err = errors.New("invalid url query params")
		response.Error(c, -1, err)
		return nil
	}

	file_h, err := c.FormFile("file")
	if err != nil {
		response.Error(c, -2, err)
		return nil
	}

	//check if upload does exist
	upload, err := services.FindUploadRecord(token, upload_id)
	if err != nil {
		response.Error(c, -3, err)
		return nil
	}

	//check if upload is expired
	is_expired, err := services.UploadExpirationCheck(upload)
	if err != nil {
		response.ServerError(c, err)
		return nil
	}
	if is_expired {
		err = errors.New("upload link is expired")
		response.Error(c, -4, err)
		return nil
	}

	//check if upload status
	if upload.Status == uploadModel.UPLOADED {
		err = errors.New("upload link was used before to upload object")
		response.Error(c, -5, err)
		return nil
	}

	// check if fileHeader can open without error
	file, err := file_h.Open()
	if err != nil {
		response.Error(c, -6, err)
		return nil
	}

	// compare file size
	if upload.Size != file_h.Size {
		err = errors.New("file size does not match")
		response.Error(c, -7, err)
		return nil
	}

	//compate file ext
	file_ext, err := helper.GetFileExtension(file_h)
	if err != nil {
		err = errors.New("could not get file's extension")
		response.Error(c, -8, err)
		return nil
	}
	if upload.Ext != file_ext {
		err = errors.New("file size extension not match")
		response.Error(c, -8, err)
		return nil
	}

	//TODO: check if file is truly the same type

	err = services.StoreFile(&file, file_h, upload)
	if err != nil {
		response.ServerError(c, err)
		return nil
	}

	response.Success(c, nil)

	return nil
}
