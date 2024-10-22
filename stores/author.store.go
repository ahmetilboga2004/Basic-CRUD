package stores

import (
	"HttpServer/models"
	"errors"
)

type AuthorStore struct {
	Authors map[int]models.Author
	NextID  int
}

// Yazar oluşturma
func (as *AuthorStore) Create(item interface{}) error {
	author, ok := item.(models.Author)
	if !ok {
		return errors.New("invalid author type")
	}
	author.ID = as.NextID
	as.Authors[as.NextID] = author
	as.NextID++
	return nil
}

func (as *AuthorStore) GetAll() (interface{}, error) {
	if len(as.Authors) == 0 {
		return nil, errors.New("no authors found")
	}
	return as.Authors, nil
}

func (as *AuthorStore) Get(id int) (interface{}, error) {
	author, exists := as.Authors[id]
	if !exists {
		return nil, errors.New("author not found")
	}
	return author, nil
}

func (as *AuthorStore) Update(id int, item interface{}) error {
	author, ok := item.(models.Author)
	if !ok {
		return errors.New("invalid author type")
	}
	if existsAuthor, exists := as.Authors[id]; exists {
		existsAuthor.Name = author.Name
		existsAuthor.Age = author.Age
		as.Authors[id] = existsAuthor
		return nil
	}
	return errors.New("author not found")
}

func (as *AuthorStore) Delete(id int) error {
	_, exists := as.Authors[id]
	if !exists {
		return errors.New("invalid author type")
	}
	delete(as.Authors, id)
	return nil
}
