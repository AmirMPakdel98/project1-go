package controllers

import (
	"c-vod/services"
	"c-vod/utils/response"
	"errors"

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
		response.Error(c, -2, err)
		return nil
	}

	data, err := services.GetFileResponseData(file)

	if err != nil {
		response.ServerError(c, err)
		return nil
	}

	response.Success(c, &data)

	return nil
}
