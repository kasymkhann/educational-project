package storage

import (
	"context"

	"REST/internal/author"
	"REST/internal/author/db/storage/model"
)

type Repository interface {
	Create(ctx context.Context, user author.Author) error
	FindAll(ctx context.Context, sortOptions model.SortOptions) (u []author.Author, err error)
	Find(ctx context.Context, id string) (author.Author, error)
	Update(ctx context.Context, users author.Author) error
	Delete(ctx context.Context, id string) error
}

type FilterOptions interface {
}
