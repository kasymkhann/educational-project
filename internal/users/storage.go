package users

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, user Users) (string, error)
	FindAll(ctx context.Context) (u []Users, err error)
	Find(ctx context.Context, id string) (Users, error)
	Update(ctx context.Context, users Users) error
	Delete(ctx context.Context, id string) error
}
