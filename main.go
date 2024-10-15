package main

import (
	"HttpServer/handlers"
	"HttpServer/stores"
	"net/http"
	"time"
)

func main() {
	bs := &stores.BookStore{}
	as := &stores.AuthorStore{}
	mux := http.NewServeMux()

	bookMux := http.NewServeMux()
	authorMux := http.NewServeMux()

	mux.Handle("/books/", http.StripPrefix("/books", bookMux))
	mux.Handle("/authors/", http.StripPrefix("/authors", authorMux))

	// Book Subrouter
	bookMux.HandleFunc("GET /", handlers.HandleGetAllBooks(bs))
	bookMux.HandleFunc("GET /{id}", handlers.HandleGetBook(bs))
	bookMux.HandleFunc("POST /", handlers.HandleCreateBook(bs))
	bookMux.HandleFunc("DELETE /{id}", handlers.HandleDeleteBook(bs))
	bookMux.HandleFunc("PUT /{id}", handlers.HandleUpdateBook(bs))

	// Author Subrouter
	authorMux.HandleFunc("POST /", handlers.HandleCreateAuthor(as))

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	srv.ListenAndServe()
}
