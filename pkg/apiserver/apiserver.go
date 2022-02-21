package apiserver

import (
	"encoding/json"
	"github.com/genridarkbkru/LinkShortenerApi/pkg"

	"github.com/genridarkbkru/LinkShortenerApi/pkg/store"
	"github.com/gorilla/mux"
	"net/http"
)

type Apiserver struct {
	store store.Store
}

type Url struct {
	Url string `json:"url"`
}

func NewServer(addr, psqlconn string) *http.Server {
	store := store.Store{}
	store.NewDB(psqlconn)
	server := Apiserver{store: store}
	router := mux.NewRouter()
	router.HandleFunc("/GetShortUrl", server.GetShortUrl).Methods("POST")
	router.HandleFunc("/GetFullUrl", server.GetFullUrl).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func (s *Apiserver) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Url
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(req.Url) > 500 || len(req.Url) == 0 {
		http.Error(w, pkg.RangeOutLenUrl.Error(), http.StatusBadRequest)
		return
	}

	var status int
	req.Url, err, status = s.store.Create(req.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(req)
}

func (s *Apiserver) GetFullUrl(w http.ResponseWriter, r *http.Request) {
	var req Url
	q := r.URL.Query()

	short_url, isHere := q["short_url"]
	if !isHere {
		http.Error(w, pkg.NotShortUrlInRequest.Error(), http.StatusBadRequest)
		return
	}

	if len(short_url[0]) > 10 || len(short_url[0]) == 0 {
		http.Error(w, pkg.RangeOutLenShortUrl.Error(), http.StatusBadRequest)
		return
	}

	var err error
	var status int
	req.Url, err, status = s.store.FindByShortUrl(short_url[0])
	if err != nil {
		http.Error(w, pkg.ErrRecordNotFound.Error(), status)
		return
	}

	//http.Redirect(w, r, req.Url, status)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(req.Url)
}
