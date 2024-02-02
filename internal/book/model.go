package book

import (
	"REST/internal/author"
)

type Book struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Authors []author.Author `json:"authors"`
}
