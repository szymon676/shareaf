package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiHandler struct {
	store Store
	addr  string
}

func NewApiHandler(store Store, addr string) *apiHandler {
	return &apiHandler{
		store: store,
		addr:  addr,
	}
}

func (ah *apiHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/paste", makeHttpHandler(ah.handleGetPaste)).Methods("GET")

	log.Print("api running on port: ", ah.addr)
	http.ListenAndServe(ah.addr, router)
}

func (ah *apiHandler) handleGetPaste(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	paste, err := ah.store.RetrievePaste(name)
	if err != nil {
		return err
	}
	return WriteJSON(200, paste, w)
}

func WriteJSON(code int, data any, w http.ResponseWriter) error {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(data)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHttpHandler(a apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			WriteJSON(404, ApiError{Error: err.Error()}, w)
		}
	}
}
