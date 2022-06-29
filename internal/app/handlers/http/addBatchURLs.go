package http

import (
	"encoding/json"
	"fmt"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto/batch"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
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

	var URLCollection []url.URL
	batchInputURLDTOCollection := []batch.BatchInputURL{}
	batchOutputURLDTOCollection := []batch.BatchOutputURL{}
	err = json.Unmarshal(bodyJSON, &batchInputURLDTOCollection)

	if err != nil {
		log.Println(`error while unmarshalling batch with URLs. ` + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, databaseURL := range batchInputURLDTOCollection {
		url, err := h.shorteningService.GetURLEntityByLongURL(databaseURL.OriginalURL)

		if err != nil || len(url.LongURL) == 0 {
			log.Println(`error while getting URL entity by long URL: ` + databaseURL.OriginalURL)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		url.UserID = h.contextStorageService.GetUserTokenFromContext(r.Context())

		h.memoryService.SaveURL(url)
		h.storageService.SaveURL(url)

		shortURLWithDomain := h.shorteningService.GetShortURLWithDomain(url.ShortURL)

		batchOutputURLDTOCollection = append(
			batchOutputURLDTOCollection, batch.NewBatchOutputURL(databaseURL.CorrelationID, shortURLWithDomain),
		)

		URLCollection = append(URLCollection, url)
	}

	h.databaseURLService.SaveBatchURLs(URLCollection)

	databaseBatchOutputURLDTOCollectionJSON, err := json.Marshal(batchOutputURLDTOCollection)

	if err != nil {
		log.Println(`error while marshalling output URL collection with length `, len(batchOutputURLDTOCollection))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("output URL collection with ", len(batchOutputURLDTOCollection), " URLs added.")

	w.WriteHeader(http.StatusCreated)
	w.Write(databaseBatchOutputURLDTOCollectionJSON)
}
