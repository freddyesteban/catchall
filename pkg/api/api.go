package api

import (
	"github.com/freddyesteban/catchall/pkg/database"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Api struct {
	router *mux.Router
	DB database.ApiDatabase
	log *logrus.Entry
}

func NewApi(db database.ApiDatabase, log *logrus.Logger) *Api {
	router := mux.NewRouter()
	logger := log.WithFi
	return &Api{
		router: router,
		DB: db,
		log: log,
	}
}

func (api *Api) Run() {
	api.log.Fatal(http.ListenAndServe(":8000", api.router))
}
