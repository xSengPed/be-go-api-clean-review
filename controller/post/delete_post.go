package postcontroller

import (
	postservice "be-nuxt-fiber/services/post"
	"be-nuxt-fiber/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type deletePostRequest struct {
	PostId string `validate:"required" query:"post_id"`
}

func DeletePost(ctx *fiber.Ctx) error {

	req := new(deletePostRequest)
	err := ctx.QueryParser(req)

	fmt.Println(req.PostId)
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

	err = postservice.DeletePost(ctx.Context(), req.PostId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post deleted"})

}
