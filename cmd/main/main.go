package main

import (
	"context"
	"github.com/freddyesteban/catchall/pkg/api"
	"github.com/freddyesteban/catchall/pkg/database/mongodb"
	"github.com/sirupsen/logrus"
	"os"
)

// TODO: allow args to be passed in for app options
func main() {
	ctx := context.Background()
	logger := createLogger()

	// TODO: Update the hard-coded uri
	db := mongodb.NewMongoDB(ctx, "fakeURI", logger)
	//err := db.Connect()
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//// It is best practice to keep a client that is connected to MongoDB around so that the application can
	//// make use of connection pooling, you don't want to open and close a connection for each query.
	//defer db.Disconnect()

	newApi := api.NewApi(db, logger)
	newApi.RegisterRoutes()
	newApi.Run()
}

func createLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	// TODO: add ability to pass this via commandline interface
	log.SetLevel(logrus.DebugLevel)

	return log
}
