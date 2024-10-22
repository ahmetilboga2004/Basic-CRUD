package models

type Book struct {
	ID     int
	Title  string `json:"title"`
	Author Author `json:"author"`
}
