package handlers

import (
	"HttpServer/stores"
	"HttpServer/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type BaseHandler[T stores.Validator[T]] struct {
	Store stores.Store[T]
}

func NewBaseHandler[T stores.Validator[T]](store stores.Store[T]) *BaseHandler[T] {
	return &BaseHandler[T]{Store: store}
}

func (h *BaseHandler[T]) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.Store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, items)
}

func (h *BaseHandler[T]) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	item, err := h.Store.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, item)
}

func (h *BaseHandler[T]) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var item T
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	result, err := h.Store.Create(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, result)
}

func (h *BaseHandler[T]) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var item T
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	err = h.Store.Update(id, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, item)
}

func (h *BaseHandler[T]) HandleDelete(w http.ResponseWriter, r *http.Request) {
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
