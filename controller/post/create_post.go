package postcontroller

import (
	"be-nuxt-fiber/datamodels"
	postservice "be-nuxt-fiber/services/post"
	"be-nuxt-fiber/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createPostReq struct {
	Topic   string `json:"topic" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func CreatePost(ctx *fiber.Ctx) error {
	req := new(createPostReq)
	err := ctx.BodyParser(req)
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

	err = postservice.CreatePost(ctx.Context(), &datamodels.Post{
		Topic:    req.Topic,
		Content:  req.Content,
		CreateAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdateAt: primitive.NewDateTimeFromTime(time.Now()),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Post created",
	})

}
