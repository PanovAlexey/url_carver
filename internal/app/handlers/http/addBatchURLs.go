package http

import (
	"encoding/json"
	"fmt"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto/database"
	"io"
	"log"
	"net/http"
)

func (h *httpHandler) HandleAddBatchURLs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	databaseBatchInputURLDTOCollection := []database.DatabaseBatchInputURL{}
	databaseBatchOutputURLDTOCollection := []database.DatabaseBatchOutputURL{}
	err = json.Unmarshal(bodyJSON, &databaseBatchInputURLDTOCollection)

	if err != nil {
		log.Println(`error while unmarshalling batch with URLs.`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, databaseURL := range databaseBatchInputURLDTOCollection {
		url, err := h.shorteningService.GetURLEntityByLongURL(databaseURL.Original_url)

		if err != nil || len(url.LongURL) == 0 {
			log.Println(`error while getting URL entity by long URL: ` + databaseURL.Original_url)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		url.SetUserToken(h.contextStorageService.GetUserTokenFromContext(r.Context()))

		// Todo to transaction
		h.memoryService.SaveURL(url)
		h.storageService.SaveURL(url)
		h.databaseURLService.SaveURL(url)

		shortURLWithDomain, err := h.shorteningService.GetShortURLWithDomain(url.GetShortURL())

		databaseBatchOutputURLDTOCollection = append(
			databaseBatchOutputURLDTOCollection, database.NewDatabaseBatchOutputURL(databaseURL.CorrelationID, shortURLWithDomain),
		)
	}

	databaseBatchOutputURLDTOCollectionJSON, err := json.Marshal(databaseBatchOutputURLDTOCollection)

	if err != nil {
		log.Println(`error while marshalling output URL collection with length `, len(databaseBatchOutputURLDTOCollection))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("output URL collection with ", len(databaseBatchOutputURLDTOCollection), " URLs added.")

	w.WriteHeader(http.StatusCreated)
	w.Write(databaseBatchOutputURLDTOCollectionJSON)

	return
}
