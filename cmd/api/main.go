package main

import (
	"flag"
	"fmt"

	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocisse/ecom-site/internal/drivers"
	"github.com/gocisse/ecom-site/internal/models"
)

const Version = "1.0.1"
const CssVersion = "5.2.0"

type config struct {
	port int
	env  string

	db struct {
		dsn string
	}

	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	DB       *models.DBModels
}

func (app *application) server() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	app.infoLog.Printf("Starting server on port %d in mode %s", app.config.port, app.config.env)

	return srv.ListenAndServe()
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 3001, "port to listen on for web requests Frontend")
	flag.StringVar(&cfg.env, "env", "development", "environment to run in  Development, Test, or Production")
	flag.StringVar(&cfg.db.dsn, "db", "mac:momo22@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "postgres dsn")

	flag.Parse()

	inforLog := log.New(os.Stdout, "//INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "//ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
	cfg.stripe.key = os.Getenv("STRIPE_KEY")

	conn, err := drivers.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  inforLog,
		errorLog: errorLog,
		version:  Version,
		DB: &models.DBModels{
			DB: conn,
		},
	}

	if err := app.server(); err != nil {
		app.errorLog.Fatal(err)
	}

}
