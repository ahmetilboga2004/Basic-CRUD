package handlers

import (
	"HttpServer/models"
	"HttpServer/stores"
	"HttpServer/utils"
	"net/http"
)

type BookHandler struct {
	*BaseHandler[*models.Book]
	bookStore *stores.BookStore
}

func NewBookHandler(store *stores.BookStore) *BookHandler {
	return &BookHandler{
		BaseHandler: NewBaseHandler(store),
		bookStore:   store,
	}
}

func (h *BookHandler) HandleFindByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	book, err := h.bookStore.FindByTitle(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, book)
}
