package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Book struct {
	ID     int
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type BookStore struct {
	books  []Book
	NextId int
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type AuthorStore struct {
	authors []Author
	NextId  int
}

func (bs *BookStore) CreateBook(title string, author Author) (Book, error) {
	book := Book{
		ID:     bs.NextId,
		Title:  title,
		Author: author,
	}
	bs.books = append(bs.books, book)
	bs.NextId++
	return book, nil
}

func (as *AuthorStore) CreateAuthor(name string, age int) (Author, error) {
	author := Author{
		ID:   as.NextId,
		Name: name,
		Age:  age,
	}

	as.authors = append(as.authors, author)
	as.NextId++
	return author, nil
}

func (bs *BookStore) GetAllBooks() ([]Book, error) {
	if len(bs.books) <= 0 {
		return nil, errors.New("Herhangi bir kitap bulunamadı")
	}

	return bs.books, nil
}

func (bs *BookStore) DeleteBook(id int) error {
	if id <= len(bs.books) {
		return errors.New("Geçersiz ID")
	}
	for i, book := range bs.books {
		if book.ID == id {
			bs.books = append(bs.books[:i], bs.books[i+1:]...)
		}
	}
	return nil
}

func (bs *BookStore) GetBook(id int) (Book, error) {
	if id < 0 || id >= len(bs.books) {
		return Book{}, errors.New("Geçersiz ID")
	}

	book := bs.books[id]
	if (book == Book{}) {
		return Book{}, errors.New("Kitap bulunamadı")
	}

	return book, nil

}

func (bs *BookStore) UpdateBook(id int, title string) error {
	for i, book := range bs.books {
		if book.ID == id {
			bs.books[i].Title = title
			return nil
		}
	}
	return errors.New("Güncellemek istediğiniz kitap bulunamadı")
}

func handleGetAllBooks(bs *BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := bs.GetAllBooks()
		if err != nil {
			http.Error(w, "Listelenecek kitap bulunamadı", http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(books)
		}
	}
}

func handleGetBook(bs *BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		fmt.Println("Params Id : ", idInt)
		book, err := bs.GetBook(idInt)
		if err != nil {
			http.Error(w, "Aradaığınız kitap bulunamadı", http.StatusNotFound)
			return
		} else {
			json.NewEncoder(w).Encode(book)
			return
		}

	}
}

func handleCreateBook(bs *BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book Book
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
func handleCreateAuthor(as *AuthorStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author Author
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields() // Beklenmeyen alanları reddet

		err := decoder.Decode(&author)
		if err != nil {
			http.Error(w, "Geçersiz veri", http.StatusBadRequest)
			return
		}

		// Alan doğrulamaları
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

func handleDeleteBook(bs *BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Geçersiz ID", http.StatusBadRequest)
			return
		}
		bs.DeleteBook(idInt)
		w.WriteHeader(http.StatusAccepted)
	}
}

func handleUpdateBook(bs *BookStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idInt, err1 := strconv.Atoi(r.PathValue("id"))
		if err1 != nil {
			http.Error(w, "Geçersiz Id", http.StatusBadRequest)
			return
		}
		var book Book
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&book)
		if err != nil {
			http.Error(w, "Geçersiz veri girdiniz", http.StatusBadRequest)
			return
		}
		err2 := bs.UpdateBook(idInt, book.Title)
		fmt.Println("Err2 : ", err2)
		if err2 != nil {
			http.Error(w, "Geçersiz veri girdiniz", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)

	}
}

func main() {
	bs := &BookStore{}
	as := &AuthorStore{}
	mux := http.NewServeMux()

	bookMux := http.NewServeMux()
	authorMux := http.NewServeMux()

	mux.Handle("/books/", http.StripPrefix("/books", bookMux))
	mux.Handle("/authors/", http.StripPrefix("/authors", authorMux))

	// Book Subrouter
	bookMux.HandleFunc("GET /", handleGetAllBooks(bs))
	bookMux.HandleFunc("GET /{id}", handleGetBook(bs))
	bookMux.HandleFunc("POST /", handleCreateBook(bs))
	bookMux.HandleFunc("DELETE /{id}", handleDeleteBook(bs))
	bookMux.HandleFunc("PUT /{id}", handleUpdateBook(bs))

	// Author Subrouter
	authorMux.HandleFunc("POST /", handleCreateAuthor(as))

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	srv.ListenAndServe()
}
