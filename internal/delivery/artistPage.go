package delivery

import (
	"net/http"
	"strconv"
	"zsandibe/pkg"
)

func (h *Handler) ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/artist/" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	artist, err := pkg.GiveData(&h.Artist, &h.Relation)
	// id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/artist/"))
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || !(id > 0 && id < len(artist.Artists)+1) {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err = h.templ.ExecuteTemplate(w, "artist.html", &artist.Artists[id-1]); err != nil {
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
