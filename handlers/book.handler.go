package handlers

import (
	"HttpServer/models"
	"HttpServer/stores"
	"HttpServer/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type BookHandler struct {
	Store *stores.BookStore
}

func NewBookHandler(store *stores.BookStore) *BookHandler {
	return &BookHandler{
		Store: store,
	}
}

func (h *BookHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	books, err := h.Store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, books)
}

func (h *BookHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	book, err := h.Store.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, book)
}

func (h *BookHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	err = h.Store.Create(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, book)
}

func (h *BookHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	err = h.Store.Update(id, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, book)
}

func (h *BookHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	err = h.Store.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
