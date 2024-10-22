package main

import (
	"HttpServer/handlers"
	"HttpServer/models"
	"HttpServer/stores"
	"net/http"
	"time"
)

func main() {
	bs := &stores.BookStore{
		Books:  make(map[int]models.Book),
		NextID: 1,
	}
	as := &stores.AuthorStore{
		Authors: make(map[int]models.Author),
		NextID:  1,
	}

	bookHandler := handlers.NewBookHandler(bs)
	authorHandler := handlers.NewAuthorHandler(as)

	mux := http.NewServeMux()

	bookMux := http.NewServeMux()
	authorMux := http.NewServeMux()

	mux.Handle("/books/", http.StripPrefix("/books", bookMux))
	mux.Handle("/authors/", http.StripPrefix("/authors", authorMux))

	// Book Subrouter
	bookMux.HandleFunc("GET /", bookHandler.HandleGetAll)
	bookMux.HandleFunc("GET /{id}", bookHandler.HandleGetByID)
	bookMux.HandleFunc("POST /", bookHandler.HandleCreate)
	bookMux.HandleFunc("PUT /{id}", bookHandler.HandleUpdate)
	bookMux.HandleFunc("DELETE /{id}", bookHandler.HandleDelete)

	// Author Subrouter
	authorMux.HandleFunc("GET /", authorHandler.HandleGetAll)
	authorMux.HandleFunc("GET /{id}", authorHandler.HandleGetByID)
	authorMux.HandleFunc("POST /", authorHandler.HandleCreate)
	authorMux.HandleFunc("PUT /{id}", authorHandler.HandleUpdate)
	authorMux.HandleFunc("DELETE /{id}", authorHandler.HandleDelete)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	srv.ListenAndServe()
}
