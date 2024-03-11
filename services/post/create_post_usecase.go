package postservice

import (
	"be-nuxt-fiber/datamodels"
	"be-nuxt-fiber/repositories"
	"context"
)

func CreatePost(ctx context.Context, post *datamodels.Post) error {
	err := repositories.CreatePost(ctx, post)
	if err != nil {
		return err
	}

	return nil
}
