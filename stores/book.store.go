package stores

import (
	"HttpServer/models"
)

type BookStore struct {
	*BaseStore[*models.Book]
}

func NewBookStore() *BookStore {
	return &BookStore{
		BaseStore: NewBaseStore[*models.Book](),
	}
}
