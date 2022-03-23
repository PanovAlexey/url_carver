package http

import (
	"log"
	"net/http"
)

func (h *httpHandler) HandlePingDatabase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := h.databaseService.CheckDatabaseAvailability()
	if err != nil {
		log.Println("error was encountered while processing the ping request: ", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	} else {
		log.Println(`database ping request succeeded`)
	}

	w.WriteHeader(http.StatusOK)
}
