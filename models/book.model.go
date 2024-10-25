package models

import "errors"

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (b *Book) Validate() error {
	if b.Title == "" || len(b.Title) > 20 {
		return errors.New("invalid book Title")
	} else if b.Desc == "" || len(b.Desc) < 20 {
		return errors.New("invalid book desc")
	}
	return nil
}
