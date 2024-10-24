package stores

import (
	"HttpServer/models"
)

type AuthorStore struct {
	*BaseStore[*models.Author]
}

func NewAuthorStore() *AuthorStore {
	return &AuthorStore{
		BaseStore: NewBaseStore[*models.Author](),
	}
}
