package handlers

import (
	"HttpServer/models"
	"HttpServer/stores"
)

type AuthorHandler struct {
	*BaseHandler[*models.Author]
}

func NewAuthorHandler(store stores.Store[*models.Author]) *AuthorHandler {
	return &AuthorHandler{
		BaseHandler: NewBaseHandler(store),
	}
}
