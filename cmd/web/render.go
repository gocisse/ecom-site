package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type TemplateData struct {
	StringMap            map[string]string
	IntMap               map[string]int
	FloatMap             map[string]float64
	Data                 map[string]interface{}
	CSRFToken            string
	API                  string
	Warning              string
	Error                string
	StripePublishableKey string
	StripeSecretKey      string
}

//functions to pass in the template
var functions = template.FuncMap{
	"formatCurrency": formatCurrency,
	"formatAmount":   formatAmount,
}

func formatCurrency(n int) string {

	f := float32(n / 100)
	return fmt.Sprintf("$%.2f", f)
}

func formatAmount(n int) string {
	f := float32(n / 10000)
	return fmt.Sprintf("$%.2f", f)
}

//go:embed templates
var templateFS embed.FS

func (app *application) addDefaultValue(td *TemplateData, r *http.Request) *TemplateData {
	td.API = app.config.api
	td.StripePublishableKey = app.config.stripe.key
	td.StripeSecretKey = app.config.stripe.secret
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *TemplateData, partials ...string) error {
	var t *template.Template
	var err error

	//template to render
	templateToRender := fmt.Sprintf("templates/%s.page.html", page)

	//check if template is in our cache
	_, templateInCache := app.templateCache[templateToRender]

	if app.config.env == "production" && templateInCache {
		t = app.templateCache[templateToRender]
	} else {
		//load template from file system
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			return err
		}

		//add template to cache
		app.templateCache[templateToRender] = t
	}

	if td == nil {
		td = &TemplateData{}
	}

	td = app.addDefaultValue(td, r)

	err = t.Execute(w, td)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) parseTemplate(partials []string, page string, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.html", x)
		}
	}

	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.html", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.html", strings.Join(partials, " "), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.html", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.html", templateToRender)
	}

	if err != nil {
		return nil, err
	}

	//set our template cache to variable t and return it
	app.templateCache[templateToRender] = t

	return t, nil
}
