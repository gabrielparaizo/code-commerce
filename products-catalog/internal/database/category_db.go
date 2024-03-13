package database

import "database/sql"

type CategoryDB struct {
	database *sql.DB
}

func NewCategoryDB(database *sql.DB) *CategoryDB {
	return &CategoryDB{database: database}
}
