package mongodb

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	Client  *mongo.Client
	URI     string
	options *options.ClientOptions
	log     *logrus.Entry
	ctx     context.Context
}

func NewMongoDB(ctx context.Context, uri string, log *logrus.Logger) (database *MongoDB) {
	logger := log.WithField("Logger", "MongoDB")
	logger.Debug("Setting up database...")
	database = &MongoDB{
		URI:     uri,
		log:     logger,
		options: options.Client().ApplyURI(uri),
		ctx:     ctx,
	}

	return
}

func (mdb *MongoDB) Connect() (err error) {
	mdb.log.Debug("Checking connection to database...")
	ctx, cancel := context.WithTimeout(mdb.ctx, 10*time.Second)
	defer cancel()

	mdb.Client, err = mongo.Connect(mdb.ctx, mdb.options)
	if err != nil {
		mdb.log.Error("Error connecting to mongodb: ", err.Error())
		return
	}
	err = mdb.Client.Ping(ctx, nil)
	return
}

func (mdb *MongoDB) Disconnect() {
	if err := mdb.Client.Disconnect(mdb.ctx); err != nil {
		mdb.log.Error("Error disconnecting from mongodb: ", err.Error())
	}
}

func (mdb *MongoDB) Query() error {
	return nil
}
