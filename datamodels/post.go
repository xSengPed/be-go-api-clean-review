package datamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Topic    string             `json:"topic" bson:"topic"`
	Content  string             `json:"content" bson:"content"`
	CreateAt primitive.DateTime `json:"create_at" bson:"create_at"`
	UpdateAt primitive.DateTime `json:"update_at" bson:"update_at"`
}
