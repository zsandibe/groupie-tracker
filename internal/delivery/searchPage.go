package delivery

// func (h *Handler) SearchHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		w.Header().Set("Allow", "GET")
// 		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 		return
// 	}
// 	text := r.FormValue("search")
// 	if !pkg.IsPrintable(text) {
// 		ErrorPage(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 		return
// 	}

// 	artist, err := pkg.SearchGroup(text)
// 	// fmt.Println(text)
// 	fmt.Println(artist.FoundArtists)
// 	if err != nil {
// 		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	if err := h.templ.ExecuteTemplate(w, "index.html", &artist); err != nil {
// 		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// }
