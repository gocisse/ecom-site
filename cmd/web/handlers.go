package main

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		app.errorLog.Printf("Error rendering template: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
