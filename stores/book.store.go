package stores

import (
	"HttpServer/models"
	"errors"
)

type BookStore struct {
	Books  map[int]models.Book
	NextID int
}

// Kitap oluşturma
func (bs *BookStore) Create(item any) error {
	book, ok := item.(models.Book)
	if !ok {
		return errors.New("invalid item type")
	}
	if book.Title == "" || len(book.Title) > 20 {
		return errors.New("invalid book title")
	}
	book.ID = bs.NextID
	bs.Books[bs.NextID] = book
	bs.NextID++
	return nil
}

// Tüm kitapları listeleme
func (bs *BookStore) GetAll() (any, error) {
	if len(bs.Books) == 0 {
		return nil, errors.New("no books found")
	}
	return bs.Books, nil
}

// Tek bir kitabı getirme
func (bs *BookStore) Get(id int) (any, error) {
	book, exists := bs.Books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

// Kitap güncelleme
func (bs *BookStore) Update(id int, item any) error {
	book, exists := item.(models.Book)
	if !exists {
		return errors.New("invalid book type")
	}
	if existsBook, exists := bs.Books[id]; exists {
		existsBook.Title = book.Title
		bs.Books[id] = existsBook
		return nil
	}
	return errors.New("book not found")
}

// Kitap silme
func (bs *BookStore) Delete(id int) error {
	_, exists := bs.Books[id]
	if !exists {
		return errors.New("book not found")
	}
	delete(bs.Books, id)
	return nil
}
