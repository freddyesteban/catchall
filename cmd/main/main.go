package main

import (
	"context"
	"github.com/freddyesteban/catchall/pkg/api"
	"github.com/freddyesteban/catchall/pkg/database"
	"github.com/sirupsen/logrus"
	"os"
)

// TODO: allow args to be passed in for app options
func main() {
	ctx := context.Background()
	log := createLogger()

	// TODO: Update the hard-coded uri
	db, err := database.NewMongoDB(ctx, "fakeURI", log)
	if err != nil {
		log.Error("unexpected error creating database connection")
		log.Fatal(err)
	}

	app := api.NewApi(db, log)
	app.Run()
}

func createLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.SetReportCaller(true)

	return log
}