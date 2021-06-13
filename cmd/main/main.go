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

	newApi := api.NewApi(db, logger)
	newApi.Run()
}

func createLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	// TODO: add ability to pass this via commandline interface
	log.SetLevel(logrus.DebugLevel)

	return log
}
