package api

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"mailgun_challenge/pkg/database"
)

type Api struct {
	Router *mux.Router
	DB database.ApiDatabase
}

func main() {
	client, err := mongo.NewClient()
}
