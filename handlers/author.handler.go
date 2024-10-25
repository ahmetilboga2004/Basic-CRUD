package handlers

import (
	"HttpServer/models"
	"HttpServer/stores"
)

type AuthorHandler struct {
	*BaseHandler[*models.Author]
	authorStore *stores.AuthorStore
}

func NewAuthorHandler(store *stores.AuthorStore) *AuthorHandler {
	return &AuthorHandler{
		BaseHandler: NewBaseHandler(store),
		authorStore: store,
	}
}
