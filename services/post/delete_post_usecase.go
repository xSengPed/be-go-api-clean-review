package postservice

import (
	"be-nuxt-fiber/repositories"
	"context"
	"fmt"
)

func DeletePost(ctx context.Context, postId string) error {
	fmt.Println(postId)
	err := repositories.DeletePost(ctx, postId)
	if err != nil {
		return err
	}
	return nil
}
