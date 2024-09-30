package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Print()
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "Internal Server Error"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	message := "Not Found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedErrorResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s Method Not Allowed", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
