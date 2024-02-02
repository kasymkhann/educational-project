package model

import (
	"REST/internal/author/db/storage"
	"REST/pkg/api/filter"
)

type filterOptions struct {
	limit  int
	fields []filter.Field
}

func NewFilterOptions(options filter.Options) storage.FilterOptions {
	return &filterOptions{limit: options.Limit(), fields: options.Fields()}
}
