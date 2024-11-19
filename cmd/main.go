package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"snip/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger        *slog.Logger
	snips         *models.SnipModel
	templateCache map[string]*template.Template
}

func main() {
	// create mySQL Pool

	// pull SQL settings from .config
	cfg := GetConfig()

	dsn := flag.String(cfg.Dialect(), cfg.ConnectionInfo(), cfg.Description)
	port := flag.String("port", ":"+cfg.ServerPort, cfg.ServerPortDesc)

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		snips:         &models.SnipModel{DB: db},
		templateCache: templateCache,
	}

	logger.Info("starting server", "port", *port)
	err = http.ListenAndServe(*port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
