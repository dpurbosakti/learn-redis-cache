package main

import "net/http"

type BookHandler struct {
}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (h *BookHandler) getBySubject(w http.ResponseWriter, r *http.Request) {

	getDTO := BookReqDTO{}

	if r.URL.Query().Get("subject") != "" {
		getDTO.Subject = r.URL.Query().Get("subject")
	}

	data, err := GetBooksBySubject(r.Context(), &getDTO)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(
		w,
		http.StatusOK,
		data,
	)
}
