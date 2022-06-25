package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

const Version = "1.0.1"
const CssVersion = "5.2.0"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}

	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 3000, "port to listen on for web requests Frontend")
	flag.StringVar(&cfg.env, "env", "development", "environment to run in  Development, Test, or Production")
	flag.StringVar(&cfg.api, "api", "http://localhost:3001", "url to the api server")
	
	flag.Parse()

	inforLog := log.New(os.Stdout, "//INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "//ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
	cfg.stripe.key = os.Getenv("STRIPE_KEY")

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       inforLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       Version,
	}

	
}
