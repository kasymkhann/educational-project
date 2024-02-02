package author

import (
	"context"
	"fmt"

	"REST/internal/author/db/storage"
	"REST/internal/author/db/storage/model"
	"REST/pkg/api/filter"
	"REST/pkg/logging"
)

type Service struct {
	repo   storage.Repository
	logger *logging.Logger
}

func NewService(repo storage.Repository, logger *logging.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetAll(ctx context.Context, filterOptions filter.Options, SortOptions model.SortOptions) ([]Author, error) {

	sortOpt := model.SortOptions{
		Filed: SortOptions.Filed,
		Order: SortOptions.Order,
	}

	u, err := s.repo.FindAll(ctx, sortOpt)
	if err != nil {
		return nil, fmt.Errorf("Error: err in FindAll  author")
	}
	return u, err
}

// TODO: Do all methods
