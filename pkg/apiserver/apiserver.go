package apiserver

import (
	"LinkShortenerApi/pkg/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Apiserver struct {
	store store.Store
}

func NewServer(addr string) *http.Server {
	server := Apiserver{store: store.Store{}}
	router := mux.NewRouter()
	router.HandleFunc("/GetShortUrl", server.GetShortUrl).Methods("POST")
	router.HandleFunc("/GetFullUrl", server.GetFullUrl).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

type Url struct {
	Url string `json:"url"`
}

func (s *Apiserver) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Url
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Реализовать логику
	req.Url = "cool"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)
}

func (s *Apiserver) GetFullUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Реализовать логику
	response := Url{
		Url: "dfdf",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
