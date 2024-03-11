package postcontroller

import (
	postservice "be-nuxt-fiber/services/post"
	"be-nuxt-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

type getPostRequest struct {
	PostId string `validate:"required" query:"post_id"`
}

func GetPost(ctx *fiber.Ctx) error {

	req := new(getPostRequest)
	err := ctx.QueryParser(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	err = utils.V.Struct(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	post, err := postservice.GetPost(ctx.Context(), req.PostId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(post)

}
