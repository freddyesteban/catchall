package api

import (
	"encoding/json"
	"fmt"
	"github.com/freddyesteban/catchall/pkg/database/mongodb"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func (a *Api) createEventsHandler(w http.ResponseWriter, r *http.Request) {
	domain := mux.Vars(r)["domain"]
	if len(strings.TrimSpace(domain)) == 0 {
		http.Error(w, fmt.Sprintf("Not a valid domain name %s", domain), http.StatusBadRequest)
		return
	}

	event := new(mongodb.Event)
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		a.log.Errorf("Error decoding request body: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := event.Validate(); err != nil {
		a.log.Errorf("Validation error: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
