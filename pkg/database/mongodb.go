package database

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	log *logrus.Entry
}


func NewMongoDB(uri string) (database *MongoDB, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}
	database = &MongoDB{
		Client: client,
		log: logrus.New().WithFields(logrus.Fields{
			"name": "mongodb",
		}),
	}

	return
}
