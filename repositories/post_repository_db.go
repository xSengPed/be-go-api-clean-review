package repositories

import (
	"be-nuxt-fiber/datamodels"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

func GetPost(ctx context.Context, postId string) (*datamodels.Post, error) {
	fmt.Println(postId)
	objectId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}

	result := DB.FindOne(ctx, filter)

	post := new(datamodels.Post)
	err = result.Decode(post)
	if err != nil {
		return nil, err

	}

	return post, nil
}

func UpdatePost(ctx context.Context, postId string) error {
	return nil
}

func DeletePost(ctx context.Context, postId string) error {

	fmt.Println(postId)
	objectId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectId}
	_, err = DB.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
