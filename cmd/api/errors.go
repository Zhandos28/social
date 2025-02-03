package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSON(w, http.StatusInternalServerError, map[string]string{
		"message": "the server encountered an unexpected error",
		"error":   err.Error(),
	})
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("bad request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSON(w, http.StatusBadRequest, map[string]string{
		"message": "the server encountered an unexpected error",
		"error":   err.Error(),
	})
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("now found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSON(w, http.StatusNotFound, map[string]string{
		"message": "the server can not find the requested resource",
		"error":   err.Error(),
	})
}

func (app *application) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("conflict error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSON(w, http.StatusConflict, map[string]string{
		"message": "the server allows you to define yours own resource",
		"error":   err.Error(),
	})
}
