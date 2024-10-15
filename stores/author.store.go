package stores

import (
	"HttpServer/models"
)

type AuthorStore struct {
	authors []models.Author
	NextId  int
}

func (as *AuthorStore) CreateAuthor(name string, age int) (models.Author, error) {
	author := models.Author{
		ID:   as.NextId,
		Name: name,
		Age:  age,
	}

	as.authors = append(as.authors, author)
	as.NextId++
	return author, nil
}
