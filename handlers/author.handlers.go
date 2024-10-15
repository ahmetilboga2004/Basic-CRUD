package handlers

import (
	"HttpServer/models"
	"HttpServer/stores"
	"encoding/json"
	"net/http"
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
