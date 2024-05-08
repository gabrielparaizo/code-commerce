package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategoty(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}
