package response

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, data any) {

	c.JSON(fiber.Map{
		"code":  0,
		"error": nil,
		"data":  data,
	})
}

func Error(c *fiber.Ctx, code int, err error) {

	c.JSON(fiber.Map{
		"code":  code,
		"error": err.Error(),
		"data":  nil,
	})
}

func Custom(c *fiber.Ctx, code int, err error, data any) {

	if err == nil {
		c.JSON(fiber.Map{
			"code":  code,
			"error": nil,
			"data":  data,
		})
	} else {
		c.JSON(fiber.Map{
			"code":  code,
			"error": err.Error(),
			"data":  data,
		})
	}
}

func ServerError(c *fiber.Ctx, err error) {

	c.JSON(fiber.Map{
		"code":  -500,
		"error": err.Error(),
		"data":  nil,
	})
}
