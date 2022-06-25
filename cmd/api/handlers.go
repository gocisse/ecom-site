package main

import (
	"fmt"
	"net/http"
)

func (app *application) apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit the api endpoint")
}
