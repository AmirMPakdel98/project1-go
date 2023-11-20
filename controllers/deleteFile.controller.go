package controllers

import (
	"c-vod/services"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func DeleteFile(c *fiber.Ctx) error {

	object_id, err := c.ParamsInt("id", -1)

	if object_id == -1 || err != nil {
		return errors.New("object_id param is missing or invalid")
	}

	err = services.DeleteFile(object_id)

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
