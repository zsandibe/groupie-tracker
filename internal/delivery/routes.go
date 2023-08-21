package delivery

import (
	"html/template"
	"net/http"
	"zsandibe/models"
)

type Handler struct {
	templ    *template.Template
	Artist   []models.Artist
	Relation models.Relation
}

func NewHandler() *Handler {
	return &Handler{
		templ: template.Must(template.ParseGlob("ui/templates/*.html")),
	}
}

func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.IndexPage)
	mux.HandleFunc("/artist", h.ArtistPage)
	mux.HandleFunc("/searched", h.SearchHandler)
	mux.HandleFunc("/filtered", h.FilterHandler)
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui/"))))
}
