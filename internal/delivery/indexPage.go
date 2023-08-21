package delivery

import (
	"fmt"
	"net/http"
	"zsandibe/pkg"
)

func (h *Handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	artist, err := pkg.GiveData(&h.Artist, &h.Relation)
	artist.FoundArtists = artist.Artists
	if err != nil {
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := h.templ.ExecuteTemplate(w, "index.html", &artist); err != nil {

		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("search")
	if !pkg.IsPrintable(text) {
		ErrorPage(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	artist, err := pkg.SearchGroup(text)
	// fmt.Println(text)
	fmt.Println(artist.FoundArtists)
	if err != nil {
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := h.templ.ExecuteTemplate(w, "index.html", &artist); err != nil {
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filtered" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	minCreationDate := r.FormValue("mindate")
	maxCreationDate := r.FormValue("maxdate")
	minFirstAlbum := r.FormValue("minalbumdate")
	maxFirstAlbum := r.FormValue("maxalbumdate")
	location := r.FormValue("location")
	r.ParseForm()
	members := r.Form["memberscount"]

	artist, err := pkg.FilterGroup(minCreationDate, maxCreationDate, minFirstAlbum, maxFirstAlbum, location, members)
	fmt.Println(artist.FoundArtists)
	if err != nil {
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := h.templ.ExecuteTemplate(w, "index.html", &artist); err != nil {
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
