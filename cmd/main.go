package main

import (
	"database/sql"
	"log"
	"net/http"
	"go-postgresql/interface/handler"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=db port=5432 user=go_postgresql password=password dbname=go_postgresql sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	app := handler.NewApp(db)
	router := app.Router()

	http.ListenAndServe(":8095", router)
}
