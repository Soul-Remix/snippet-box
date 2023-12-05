package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Soul-Remix/snippet-box/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type dbContext struct {
	snippets *models.SnippetModel
}

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	dbContext     *dbContext
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "snippet:pass@/snippetbox?parseTime=true", "MySQl data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: templateCache,
		dbContext: &dbContext{
			snippets: &models.SnippetModel{DB: db},
		},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Print("Starting server on :4000")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
