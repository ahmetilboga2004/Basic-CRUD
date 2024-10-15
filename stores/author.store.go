package stores

import (
	"HttpServer/models"
	"errors"
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

func (as *AuthorStore) GetAuthor(id int) (models.Author, error) {
	if id < 0 || id > len(as.authors) {
		return models.Author{}, errors.New("geçersiz ID")
	}
	author := as.authors[id]
	if author == (models.Author{}) {
		return models.Author{}, errors.New("yazar bulunamadı")
	}
	return author, nil
}

func (as *AuthorStore) GetAllAuthor() ([]models.Author, error) {
	if len(as.authors) <= 0 {
		return nil, errors.New("herhangi bir yazar bulunamadı")
	}
	return as.authors, nil
}

func (as *AuthorStore) DeleteAuthor(id int) error {
	if id < 0 || id > len(as.authors) {
		return errors.New("gerçersiz ID")
	}
	as.authors = append(as.authors[:id], as.authors[id+1:]...)
	return nil
}

func (as *AuthorStore) UpdateAuthor(id int, name string, age int) error {
	if id < 0 || id > len(as.authors) {
		return errors.New("geçersiz ID")
	}
	author := as.authors[id]
	if author == (models.Author{}) {
		return errors.New("yazar bulunamadı")
	}
	author.Name = name
	author.Age = age
	return nil
}
