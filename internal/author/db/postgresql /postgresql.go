package postgresql

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"

	"REST/internal/author"
	"REST/internal/author/db/storage/model"
	"REST/pkg/client/postgresql"
	"REST/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) *repository {
	return &repository{
		client: client,
		logger: logger,
	}

}

func (r *repository) Create(ctx context.Context, author author.Author) error {
	query := `INSERT INTO author(name) VALUES ($1)`

	if err := r.client.QueryRow(ctx, query, author.Name).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newError := fmt.Errorf("SQL Error: %s, Code: %s, Detail: %s, Where: %s, SQLState: %s", pgErr.Error(), pgErr.Code, pgErr.Detail, pgErr.Where, pgErr.SQLState())
			r.logger.Error(newError)
			return newError
		}
		return err
	}
	return nil

}
func (r *repository) FindAll(ctx context.Context, sortOptions model.SortOptions) (u []author.Author, err error) {
	qb := sq.Select("id, name, age, is_alive, created_at").From("author") // pub lic.author
	if sortOptions.Filed != "" && sortOptions.Order != "" {
		qb = qb.OrderBy(fmt.Sprintf("%s %s", sortOptions.Filed, sortOptions.Order))
	}

	sql, i, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.client.Query(ctx, sql, i...)
	if err != nil {
		return nil, err
	}

	authors := make([]author.Author, 0)

	for rows.Next() {
		var ath author.Author

		if err = rows.Scan(&ath.ID, &ath.Name, &ath.Age, &ath.IsAlive, &ath.CreatedAt); err != nil {
			return nil, err
		}
		authors = append(authors, ath)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil

}
func (r *repository) Find(ctx context.Context, id string) (author.Author, error) {
	query := `SELECT id FROM author WHERE id =  $1`

	var ath author.Author
	err := r.client.QueryRow(ctx, query, id).Scan(&ath.ID, &ath.Name)
	if err != nil {
		return author.Author{}, err
	}
	return ath, nil

}
func (r *repository) Update(ctx context.Context, user author.Author) error {
	panic("implement me")
}
func (r *repository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
