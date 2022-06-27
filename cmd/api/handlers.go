package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gocisse/ecom-site/internal/cards"
)

type apiPayload struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      string `json:"id"`
}

func (app *application) apiHandler(w http.ResponseWriter, r *http.Request) {

	var paylod apiPayload

	err := json.NewDecoder(r.Body).Decode(&paylod)
	if err != nil {
		app.errorLog.Printf("Error decoding JSON: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	okay := true

	card := cards.Card{
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.key,
		Currency: paylod.Currency,
	}

	amount, err := strconv.Atoi(paylod.Amount)
	if err != nil {
		app.errorLog.Printf("Error converting amount: %s", err)
	}

	pi, msg, err := card.Charge(amount, paylod.Currency)
	if err != nil {
		app.errorLog.Printf("Error charging card: %s", err)
		okay = false
	}
	if okay {

		out, err := json.MarshalIndent(pi, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)

	} else {
		var j = jsonResponse{
			OK:      false,
			Message: msg,
		}

		out, err := json.MarshalIndent(j, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}

}

func (app *application) GetWidgetByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	widgetID, _ := strconv.Atoi(id)

	widget, err := app.DB.GetWidgetByID(widgetID)
	if err != nil {
		app.errorLog.Printf("Error getting widget: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	out, err := json.MarshalIndent(widget, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}
