package repositories

import (
	"be-nuxt-fiber/datamodels"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(ctx context.Context, post *datamodels.Post) error {
	objectId := primitive.NewObjectID()
	post.ID = objectId

	_, err := DB.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	return nil
}

func GetPost() {

}

func UpdatePost() {

}

func DeletePost() {

}
