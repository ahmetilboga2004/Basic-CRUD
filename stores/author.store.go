package stores

import (
	"HttpServer/models"
	"database/sql"
	"errors"
)

type AuthorStore struct {
	DB *sql.DB
}

func NewAuthorStore(db *sql.DB) *AuthorStore {
	return &AuthorStore{
		DB: db,
	}
}

func (as *AuthorStore) Create(author *models.Author) error {
	_, err := as.DB.Exec("INSERT INTO authors (firstName, lastName, age) VALUES (?, ?, ?)", author.FirstName, author.LastName, author.Age)
	return err
}

func (as *AuthorStore) GetAll() ([]*models.Author, error) {
	rows, err := as.DB.Query("SELECT * FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*models.Author
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.FirstName, &author.LastName, &author.Age); err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}
	return authors, nil
}

func (as *AuthorStore) Get(id int) (*models.Author, error) {
	row := as.DB.QueryRow("SELECT * FROM authors WHERE id = ?", id)
	var author models.Author
	if err := row.Scan(&author.ID, &author.FirstName, &author.LastName, &author.Age); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("author not found")
		}
		return nil, err
	}
	return &author, nil
}

func (as *AuthorStore) Update(id int, author *models.Author) error {
	_, err := as.DB.Exec("UPDATE authors SET firstName = ?, lastName = ?, age = ? WHERE id = ?", author.FirstName, author.LastName, author.Age, id)
	return err
}

func (as *AuthorStore) Delete(id int) error {
	_, err := as.DB.Exec("DELETE FROM authors WHERE id = ?", id)
	return err
}

func (as *AuthorStore) FindByName(firstName string) (*models.Author, error) {
	row := as.DB.QueryRow("SELECT * FROM authors WHERE firstName = ?", firstName)
	var author models.Author
	err := row.Scan(&author.ID, &author.FirstName, &author.LastName, &author.Age)
	if err != nil {
		return nil, err
	}
	return &author, nil
}
