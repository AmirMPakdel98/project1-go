package controllers

import (
	"c-vod/services"
	"c-vod/utils/types"

	"github.com/gofiber/fiber/v2"
)

func CreateFile(c *fiber.Ctx) error {

	body := &types.CreateFilesReq{}

	err := c.BodyParser(body)

	if err != nil {

		c.JSON(fiber.Map{
			"code":  -1,
			"error": err.Error(),
			"data":  nil,
		})
		return nil
	}

	upload, err := services.CreateUpload(body)

	if err != nil {
		return err
	}

	upload_url := services.GetUploadURL(upload)

	c.JSON(fiber.Map{
		"code":  0,
		"error": nil,
		"data": &types.CreateFilesResData{
			Object_id:  upload.Id,
			Upload_url: upload_url,
		},
	})

	return nil
}
