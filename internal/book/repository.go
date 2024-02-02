package book

import (
	"golang.org/x/net/context"
)

type Repository interface {
	Create(ctx context.Context, book Book) error
	FindAll(ctx context.Context) (u []Book, err error)
	Find(ctx context.Context, id string) (Book, error)
	Update(ctx context.Context, book Book) error
	Delete(ctx context.Context, id string) error
}
