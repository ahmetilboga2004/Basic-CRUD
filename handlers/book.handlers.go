package handlers

import (
	"HttpServer/models"
	"HttpServer/stores"
	"encoding/json"
	"net/http"
	"strconv"
)

// GET All Books
func HandleGetAllBooks(bs *stores.BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := bs.GetAllBooks()
		if err != nil {
			http.Error(w, "Listelenecek kitap bulunamadı", http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(books)
		}
	}
}

// GET Book by ID
func HandleGetBook(bs *stores.BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Geçersiz ID", http.StatusBadRequest)
			return
		}
		book, err := bs.GetBook(idInt)
		if err != nil {
			http.Error(w, "Aradığınız kitap bulunamadı", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(book)
	}
}

// CREATE Book
func HandleCreateBook(bs *stores.BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, "Geçersiz veri", http.StatusBadRequest)
			return
		}
		createdBook, err := bs.CreateBook(book.Title, book.Author)
		if err != nil {
			http.Error(w, "Kitap oluşturulamadı", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdBook)
	}
}

// UPDATE Book
func HandleUpdateBook(bs *stores.BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Geçersiz ID", http.StatusBadRequest)
			return
		}
		var book models.Book
		err = json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, "Geçersiz veri", http.StatusBadRequest)
			return
		}
		err = bs.UpdateBook(idInt, book.Title)
		if err != nil {
			http.Error(w, "Kitap güncellenemedi", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

// DELETE Book
func HandleDeleteBook(bs *stores.BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Geçersiz ID", http.StatusBadRequest)
			return
		}
		err = bs.DeleteBook(idInt)
		if err != nil {
			http.Error(w, "Kitap silinemedi", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
