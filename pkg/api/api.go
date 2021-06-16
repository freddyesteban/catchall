package api

import (
	"fmt"
	"github.com/freddyesteban/catchall/pkg/database"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Api struct {
	router *mux.Router
	DB     database.ApiDatabase
	log    *logrus.Entry
}

func NewApi(db database.ApiDatabase, log *logrus.Logger) *Api {
	router := mux.NewRouter()
	logger := log.WithField("logger", "Api")
	logger.Debug("Setting up api...")
	return &Api{
		router: router,
		DB:     db,
		log:    logger,
	}
}

func (a *Api) RegisterRoutes() {
	a.router.HandleFunc("/events/{domain}/delivered", a.createEventsHandler).Methods(http.MethodPut)
}

func (a *Api) Run() {
	port := "8000"
	a.log.Infof("Starting server on port %s", port)
	a.log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), a.router))
}
