package controllers

import (
	"c-vod/models/uploadModel"
	"c-vod/services"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {

	token := c.Query("t", "")
	upload_id := c.Query("i", "")

	if token == "" || upload_id == "" {
		c.JSON(fiber.Map{
			"code":  -1,
			"error": "invalid url query params",
			"data":  nil,
		})
		return nil
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(fiber.Map{
			"code":  -2,
			"error": err.Error(),
			"data":  nil,
		})
		return nil
	}

	upload, err := services.FindUploadRecord(token, upload_id)

	if err != nil {
		return err
	}

	is_expired, err := services.UploadExpirationCheck(upload)

	if err != nil {
		return err
	}

	if is_expired {
		c.JSON(fiber.Map{
			"code":  -3,
			"error": "upload link is expired",
			"data":  nil,
		})
		return nil
	}

	//check if upload status
	if upload.Status == uploadModel.UPLOADED {
		c.JSON(fiber.Map{
			"code":  -4,
			"error": "file was already uploaded",
			"data":  nil,
		})
		return nil
	}

	err = services.StoreFile(file, upload)

	if err != nil {
		return err
	}

	c.JSON(fiber.Map{
		"code":  0,
		"error": nil,
		"data":  nil,
	})

	return nil
}
