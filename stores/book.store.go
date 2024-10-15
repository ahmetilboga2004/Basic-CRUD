package stores

import (
	"HttpServer/models"
	"errors"
)

type BookStore struct {
	Books  []models.Book
	NextID int
}

// Kitap oluşturma
func (bs *BookStore) CreateBook(title string, author models.Author) (models.Book, error) {
	book := models.Book{
		ID:     bs.NextID,
		Title:  title,
		Author: author,
	}
	bs.Books = append(bs.Books, book)
	bs.NextID++
	return book, nil
}

// Tüm kitapları listeleme
func (bs *BookStore) GetAllBooks() ([]models.Book, error) {
	if len(bs.Books) == 0 {
		return nil, errors.New("herhangi bir kitap bulunamadı")
	}
	return bs.Books, nil
}

// Tek bir kitabı getirme
func (bs *BookStore) GetBook(id int) (models.Book, error) {
	if id < 0 || id >= len(bs.Books) {
		return models.Book{}, errors.New("geçersiz ID")
	}
	book := bs.Books[id]
	if book == (models.Book{}) {
		return models.Book{}, errors.New("kitap bulunamadı")
	}
	return book, nil
}

// Kitap güncelleme
func (bs *BookStore) UpdateBook(id int, title string) error {
	for i, book := range bs.Books {
		if book.ID == id {
			bs.Books[i].Title = title
			return nil
		}
	}
	return errors.New("güncellemek istediğiniz kitap bulunamadı")
}

// Kitap silme
func (bs *BookStore) DeleteBook(id int) error {
	if id < 0 || id >= len(bs.Books) {
		return errors.New("geçersiz ID")
	}
	for i, book := range bs.Books {
		if book.ID == id {
			bs.Books = append(bs.Books[:i], bs.Books[i+1:]...)
			return nil
		}
	}
	return errors.New("kitap bulunamadı")
}
