package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/n0ks/snipbox/pkg/models/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	dsn := flag.String(
		"dsn",
		"web:password@/snippetbox?parseTime=true",
		"MySQL data source name testing",
	)

	flag.Parse()

	f, err := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0o666)

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Print("Should i codesign?")

	if err != nil {
		errorLog.Fatal(err)
	}

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer f.Close()
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)

	err = server.ListenAndServe()

	errorLog.Fatal(err)
}
