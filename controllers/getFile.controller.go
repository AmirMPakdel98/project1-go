package controllers

import (
	"c-vod/services"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetFile(c *fiber.Ctx) error {

	object_id, err := c.ParamsInt("id", -1)

	if object_id == -1 || err != nil {
		return errors.New("object_id param is missing or invalid")
	}

	response, err := services.GetFile(object_id)

	if err != nil {
		return err
	}

	c.JSON(fiber.Map{
		"code":  0,
		"error": nil,
		"data":  &response,
	})

	return nil
}
