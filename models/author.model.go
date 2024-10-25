package models

import "errors"

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func (a Author) Validate() error {
	if a.FirstName == "" || a.LastName == "" {
		return errors.New("invalid author name")
	}
	if a.Age <= 0 {
		return errors.New("invalid age")
	}
	return nil
}

func (a *Author) SetID(id int) {
	a.ID = id
}
