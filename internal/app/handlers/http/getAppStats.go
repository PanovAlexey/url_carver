package http

import (
	"encoding/json"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"net/http"
)

func (h *httpHandler) HandleGetAppStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usersCount, err := h.databaseUserService.GetAllUsersCount()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	appStat := dto.GetAppStatByURLsCountAndUsersCount(h.memoryService.GetAllURLsCount(), usersCount)
	appStatJSON, err := json.Marshal(appStat)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(appStatJSON)
}
