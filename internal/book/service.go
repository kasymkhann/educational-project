package book

import (
	"context"
	"fmt"

	"REST/pkg/logging"
)

type Service struct {
	repo   Repository
	logger *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *Service {
	return &Service{repo: repository, logger: logger}
}

func (s *Service) GetAll(ctx context.Context) ([]Book, error) {
	u, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error: err in FindAll book")
	}
	return u, nil
}
