package database

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	Client *mongo.Client
	URI string
	options *options.ClientOptions
	log *logrus.Logger
	ctx context.Context
	cancel context.CancelFunc
}


func NewMongoDB(ctx context.Context, uri string, log *logrus.Logger) (*MongoDB, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	database := &MongoDB{
		URI: uri,
		log: log,
		options: options.Client().ApplyURI(uri),
		ctx: timeoutCtx,
		cancel: cancel,
	}

	return database, nil
}

func (mdb *MongoDB) Connect() (err error) {
	mdb.Client, err = mongo.Connect(mdb.ctx,mdb.options)
	if err != nil {
		mdb.log.Error("unexpected error while connecting to database")
	}
	return
}

func (mdb *MongoDB) Query() error {
	return nil
}
