package stores

import (
	"HttpServer/models"
	"database/sql"
	"errors"
)

type BookStore struct {
	DB *sql.DB
}

func NewBookStore(db *sql.DB) *BookStore {
	return &BookStore{
		DB: db,
	}
}

func (bs *BookStore) Create(book *models.Book) (*models.Book, error) {
	query := `INSERT INTO books (title, desc) VALUES (?, ?) RETURNING *`
	err := bs.DB.QueryRow(query, book.Title, book.Desc).Scan(&book.ID, &book.Title, &book.Desc)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (bs *BookStore) GetAll() ([]*models.Book, error) {
	rows, err := bs.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Desc); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (bs *BookStore) Get(id int) (*models.Book, error) {
	row := bs.DB.QueryRow("SELECT * FROM books WHERE id = ?", id)
	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.Desc); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

func (bs *BookStore) Update(id int, book *models.Book) error {
	_, err := bs.DB.Exec("UPDATE books SET title = ?, desc = ? WHERE id = ?", book.Title, book.Desc, id)
	return err
}

func (bs *BookStore) Delete(id int) error {
	_, err := bs.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}

func (bs *BookStore) FindByTitle(title string) (*models.Book, error) {
	row := bs.DB.QueryRow("SELECT id, title, desc FROM books WHERE  title LIKE ?", "%"+title+"%")
	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.Desc)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
