package handlers

import (
	"HttpServer/models"
	"HttpServer/stores"
	"encoding/json"
	"net/http"
	"strconv"
)

func HandleCreateAuthor(as *stores.AuthorStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&author)
		if err != nil {
			http.Error(w, "Geçersiz veri", http.StatusBadRequest)
			return
		}

		if author.Name == "" || author.Age <= 0 {
			http.Error(w, "Geçersiz yazar bilgileri", http.StatusBadRequest)
			return
		}

		createdAuthor, err := as.CreateAuthor(author.Name, author.Age)
		if err != nil {
			http.Error(w, "Yazar oluşturulamadı", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdAuthor)
	}
}

func HandleGetAuthor(as *stores.AuthorStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "geçersiz id girdiniz", http.StatusBadRequest)
			return
		}
		author, err2 := as.GetAuthor(idInt)
		if err2 != nil {
			http.Error(w, "aradaığınız yazar bulunamadı", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(author)
	}
}

func HandleGetAllAuthors(as *stores.AuthorStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := as.GetAllAuthor()
		if err != nil {
			http.Error(w, "listelenecek yazar bulunamadı", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(authors)
	}
}

func HandleUpdateAuthor(as *stores.AuthorStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "güncellemek istediğiniz yazar bulunamadı", http.StatusNotFound)
			return
		}
		var author models.Author
		err2 := json.NewDecoder(r.Body).Decode(&author)
		if err2 != nil {
			http.Error(w, "geçersiz veri", http.StatusBadRequest)
			return
		}
		err3 := as.UpdateAuthor(idInt, author.Name, author.Age)
		if err3 != nil {
			http.Error(w, "yazar bilgileri güncellenemedi", http.StatusNotModified)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(author)
	}
}

func HandleDeleteAuthor(as *stores.AuthorStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "geçersiz ID", http.StatusBadRequest)
			return
		}
		err2 := as.DeleteAuthor(idInt)
		if err2 != nil {
			http.Error(w, "yazar silinemedi", http.StatusNotModified)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
