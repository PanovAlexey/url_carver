package http

import (
	"net/http"
)

func (h *httpHandler) HandleGetURLsByUserId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
