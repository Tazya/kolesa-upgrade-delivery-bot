package delivery

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", h.handleHealth)
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/health" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	res := StatusResponse{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(res)
}
