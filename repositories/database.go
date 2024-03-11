package repositories

import (
	"be-nuxt-fiber/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slog"
)

var DB *mongo.Collection

func ConnectToMongoDB(mongoConfig *config.MongoConfig) {
	db, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://"+mongoConfig.Username+":"+mongoConfig.Password+"@"+mongoConfig.ClusterName))
	if err != nil {
		panic(err)
	}

	slog.Info("Connected to MongoDB")
	DB = db.Database("nuxt_app").Collection("post")
}
