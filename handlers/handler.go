package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type IHandler interface {
	Handle(router chi.Router)
}

// Ok : HTTP 200 response
func Ok(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(d)
}

// Created : HTTP 201 response
func Created(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}

// Accepted : HTTP 202 Accepted
func Accepted(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

// NoContent : HTTP 204 response
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}

// BadRequest : HTTP 400 response
func BadRequest(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(d)
}

// NotFound : HTTP 404 response
func NotFound(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusNotFound)
	if d == nil {
		w.Write([]byte(`{"message" : "requested data is not found"}`))
	} else {
		json.NewEncoder(w).Encode(d)
	}
}

// InternalServerError : HTTP 500
func InternalServerError(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(d)
}
