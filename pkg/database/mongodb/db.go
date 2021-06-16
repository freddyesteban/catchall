package mongodb

import (
	"context"
	"fmt"
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
	dbName  string
}

func NewMongoDB(ctx context.Context, uri string, log *logrus.Logger) (database *MongoDB) {
	logger := log.WithField("Logger", "MongoDB")
	logger.Debug("Setting up database...")
	database = &MongoDB{
		URI:     uri,
		log:     logger,
		options: options.Client().ApplyURI(uri),
		ctx:     ctx,
		dbName:  "catchall",
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
		return err
	}
	err = mdb.Client.Ping(ctx, nil)
	return nil
}

func (mdb *MongoDB) Disconnect() {
	if err := mdb.Client.Disconnect(mdb.ctx); err != nil {
		mdb.log.Error("Error disconnecting from mongodb: ", err.Error())
	}
}

func (mdb *MongoDB) Query() error {
	return nil
}

func (mdb *MongoDB) InsertEvent(event Event) (err error) {
	mdb.log.Debug(fmt.Sprintf("Inserting event %v", event))
	ctx, cancel := context.WithTimeout(mdb.ctx, 10*time.Second)
	defer cancel()

	collection := mdb.eventsCollection()
	result, err := collection.InsertOne(ctx, event)
	if err != nil {
		mdb.log.Error("Error inserting event: ", err.Error())
		return err
	}
	mdb.log.Info("Inserted single event document: ", result.InsertedID)
	return nil
}

func (mdb *MongoDB) eventsCollection() *mongo.Collection {
	return mdb.Client.Database(mdb.dbName).Collection("events")
}
