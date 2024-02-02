package db

import (
	"context"

	"REST/internal/book"
	"REST/pkg/client/postgresql"
	"REST/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) *repository {
	return &repository{client: client, logger: logger}
}

func (repo *repository) Create(ctx context.Context, book book.Book) error {
	//TODO this is stub, реализовать логику repository
	return nil
}
func (repo *repository) FindAll(ctx context.Context) (u []book.Book, err error) {
	//TODO this is stub, реализовать логику repository
	return u, nil
}
func (repo *repository) Find(ctx context.Context, id string) (u book.Book, err error) {
	//TODO this is stub, реализовать логику repository
	return u, err
}
func (repo *repository) Update(ctx context.Context, book book.Book) error {
	//TODO this is stub, реализовать логику repository
	return nil
}
func (repo *repository) Delete(ctx context.Context, id string) error {
	//TODO this is stub, реализовать логику repository
	return nil
}
