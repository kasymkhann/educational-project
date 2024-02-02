package users

import (
	"context"

	"REST/pkg/logging"
)

type Service struct {
	Storage Storage
	Logger  logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUsers) (u Users, err error) {
	//TODO this is stub, реализовать логику
	return
}
func (s *Service) FindAll(ctx context.Context) (u []Users, err error) {
	//TODO this is stub, реализовать логику
	return u, nil
}
func (s *Service) Find(ctx context.Context, id string) (u Users, err error) {
	//TODO this is stub, реализовать логику

	return u, nil
}
func (s *Service) Update(ctx context.Context, users Users) error {
	//TODO this is stub, реализовать логику
	return nil
}
func (s *Service) Delete(ctx context.Context, id string) error {
	//TODO this is stub, реализовать логику
	return nil
}
