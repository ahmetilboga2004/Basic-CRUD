package handlers

import (
	"HttpServer/models"
	"HttpServer/stores"
)

type BookHandler struct {
	*BaseHandler[*models.Book]
}

func NewBookHandler(store stores.Store[*models.Book]) *BookHandler {
	return &BookHandler{
		BaseHandler: NewBaseHandler(store),
	}
}
