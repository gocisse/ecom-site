package main

import (
	"net/http"
	"strconv"

	"github.com/gocisse/ecom-site/internal/models"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", &TemplateData{}, "stripe-js"); err != nil {
		app.errorLog.Printf("Error rendering template: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) Succeeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Printf("Error parsing form: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	cardHolder := r.Form.Get("cardholder_name")
	email := r.Form.Get("cardholder_email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]interface{})
	data["cardHolder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	prices, _ := strconv.Atoi(paymentAmount)
	data["price"] = prices

	if err := app.renderTemplate(w, r, "succeeded", &TemplateData{
		Data: data,
	}); err != nil {
		app.errorLog.Printf("Error rendering template: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) BuyOnce(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.errorLog.Printf("Error parsing form: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	widget := models.Widgets{
		ID:          1,
		Name:        "Widget",
		Price:       1000,
		Description: "This is a very nice widget",
	}

	data := make(map[string]interface{})

	data["widget"] = widget

	if err := app.renderTemplate(w, r, "buy-once", &TemplateData{
		Data: data,
	}, "stripe-js"); err != nil {
		app.errorLog.Printf("Error rendering template: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
