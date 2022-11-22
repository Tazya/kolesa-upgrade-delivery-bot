package delivery

import (
	"encoding/json"
	"net/http"

	"kolesa-upgrade-team/delivery-bot/usecase"
)

type StatusResponse struct {
	Status string `json:"status"`
	Error string `json:"error,omitempty"` 
}


type SendAllRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Sender interface {
	SendAll(msg SendAllRequest) error
}

type Handler struct {
	sender usecase.Sender
}

func NewHandler(sender usecase.Sender) *Handler {
	return &Handler{sender: sender}
}

func (h *Handler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", h.handleHealth)
	mux.HandleFunc("/messages/sendAll", h.SendAll)
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	res := StatusResponse{
		Status: "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		res.Status = "error"
		res.Error = "Method Not Allowed"
		json.NewEncoder(w).Encode(res)
		return
	}

	if r.URL.Path != "/health" {
		res.Status = "error"
		res.Error = "Status Not Found"
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *Handler) SendAll(w http.ResponseWriter, r *http.Request) {
	res := StatusResponse{
		Status: "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		res.Status = "error"
		res.Error = "Method Not Allowed"
		json.NewEncoder(w).Encode(res)
		return
	}
	var reqBody usecase.Message
	
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		res.Status = "error"
		res.Error = "Bad Request"
		json.NewEncoder(w).Encode(res)
		return
	}
	defer r.Body.Close()

	if reqBody.Body == "" {
		res.Status = "error"
		res.Error = "Bad Request. Message must have body"
		json.NewEncoder(w).Encode(res)
		return
	}

	if err := h.sender.SendAll(reqBody); err != nil {

		res.Status = "error"
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	

	json.NewEncoder(w).Encode(res)
}
