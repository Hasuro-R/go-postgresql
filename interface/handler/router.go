package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct{}

func NewApp(db *sql.DB) *App {
	return &App{}
}

func (a *App) Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Recuord!"))
	})

	return router
}
