package controllers

import (
	"c-vod/services"
	"c-vod/utils/response"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetFile(c *fiber.Ctx) error {

	object_id, err := c.ParamsInt("id", -1)

	if object_id == -1 || err != nil {
		err = errors.New("object_id param is missing or invalid")
		response.Error(c, -1, err)
		return nil
	}

	file, err := services.FindFileRecord(object_id)

	if err != nil {

		//check if upload exists
		upload, err := services.FindUploadRecordById(object_id)

		if err != nil {
			response.Error(c, -2, err)
			return nil
		}

		//check if upload is expired
		if upload.ExpiresAt < time.Now().UnixMilli() {
			response.Error(c, -2, errors.New("upload url is expired"))
			return nil
		} else {
			response.Custom(c, 1, nil, "waiting for user to upload the object")
			return nil
		}
	}

	data, err := services.GetFileResponseData(file)

	if err != nil {
		response.ServerError(c, err)
		return nil
	}

	response.Success(c, &data)

	return nil
}
