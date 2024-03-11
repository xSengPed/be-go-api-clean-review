package postservice

import (
	"be-nuxt-fiber/datamodels"
	"be-nuxt-fiber/repositories"
	"context"
	"fmt"
)

func GetPost(ctx context.Context, postId string) (*datamodels.Post, error) {
	fmt.Println(postId)
	post, err := repositories.GetPost(ctx, postId)
	if err != nil {
		return nil, err
	}

	return post, nil
}
