package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	writeJSON(w, http.StatusInternalServerError, map[string]string{
		"message": "the server encountered an unexpected error",
		"error":   err.Error(),
	})
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	writeJSON(w, http.StatusBadRequest, map[string]string{
		"message": "the server encountered an unexpected error",
		"error":   err.Error(),
	})
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	writeJSON(w, http.StatusNotFound, map[string]string{
		"message": "the server can not find the requested resource",
		"error":   err.Error(),
	})
}

func (app *application) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("conflict error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	writeJSON(w, http.StatusConflict, map[string]string{
		"message": "the server allows you to define yours own resource",
		"error":   err.Error(),
	})
}
