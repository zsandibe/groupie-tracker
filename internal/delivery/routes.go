package delivery

import (
	"html/template"
	"net/http"
	"zsandibe/models"
)

type Handler struct {
	templ    *template.Template
	Artists  []models.Artist
	Relation models.Relation
}

func NewHandler() *Handler {
	return &Handler{
		templ: template.Must(template.ParseFiles("./ui/templates/*.html")),
	}
}

func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.IndexPage)
	mux.HandleFunc("/artist/", h.ArtistPage)
}
