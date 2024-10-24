package handlers

import (
	"HttpServer/stores"
	"HttpServer/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

// BaseHandler, genel CRUD işlemleri için temel bir yapı sağlar.
type BaseHandler[T stores.Validator[T]] struct {
	Store stores.Store[T] // Store arayüzü burada kullanılacak
}

func NewBaseHandler[T stores.Validator[T]](store stores.Store[T]) *BaseHandler[T] {
	return &BaseHandler[T]{Store: store}
}

// HandleGetAll, tüm öğeleri döndürür.
func (h *BaseHandler[T]) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.Store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, items)
}

// HandleGetByID, belirli bir ID'ye sahip öğeyi döndürür.
func (h *BaseHandler[T]) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // Burada istediğin kodu kullanıyoruz
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

// HandleCreate, yeni bir öğe ekler.
func (h *BaseHandler[T]) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var item T
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	err = h.Store.Create(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, item)
}

// HandleUpdate, belirli bir ID'ye sahip öğeyi günceller.
func (h *BaseHandler[T]) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // Burada istediğin kodu kullanıyoruz
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

// HandleDelete, belirli bir ID'ye sahip öğeyi siler.
func (h *BaseHandler[T]) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // Burada istediğin kodu kullanıyoruz
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
