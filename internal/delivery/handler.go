package delivery

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

type SendAllRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Sender interface {
	SendAll(msg SendAllRequest) error
}

type Handler struct {
	sender Sender
}

func NewHandler(sender Sender) *Handler {
	return &Handler{sender: sender}
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

func (h *Handler) SendAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var reqBody SendAllRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if reqBody.Body == "" {
		http.Error(w, "Bad Request. Message must have body", http.StatusBadRequest)
		return
	}

	if err := h.sender.SendAll(reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := StatusResponse{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(res)

}
