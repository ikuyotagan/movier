package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/ikuyotagan/movier/internal/database"
	"github.com/ikuyotagan/movier/internal/models"
)

const (
	genresTable = "genres"
)

// GenresRepository имплементация
type GenresRepository struct {
	db QueryExecutor
}

// NewGenresRepository конструктор
func NewGenresRepository(db QueryExecutor) database.GenresRepository {
	return &GenresRepository{db: db}
}

// Create genres
func (l *GenresRepository) Create(ctx context.Context, genres []*models.Genre) ([]uint64, error) {
	if len(genres) == 0 {
		return nil, database.ErrEmptyStruct
	}
	qb := sq.
		Insert(genresTable).
		Columns(
			"name",
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)
	for _, genre := range genres {
		qb = qb.Values(
			genre.Code,
		)
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := l.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	var ids []uint64
	for rows.Next() {
		id := uint64(0)
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

// All get genres
func (l *GenresRepository) All(ctx context.Context, filter *models.GenresFilter) ([]*models.Genre, error) {
	qb := sq.Select("id", "name").
		From(genresTable).
		OrderBy("id").
		PlaceholderFormat(sq.Dollar)

	result := make([]*models.Genre, 0)
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := l.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		id := uint64(0)
		name := ""
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		result = append(result, &models.Genre{
			Id:   id,
			Code: name,
		})
	}
	return result, nil
}

// All get genres
func (l *GenresRepository) Count(ctx context.Context, filter *models.GenresFilter) (uint64, error) {
	return 0, nil
}

// Update update layer
func (l *GenresRepository) Update(ctx context.Context, layer *models.Genre) (*models.Genre, error) {
	return nil, nil
}

// Delete delete layer
func (l *GenresRepository) Delete(ctx context.Context, genreIDs []uint64) error {
	qb := sq.Delete(genresTable).
		Where(sq.Eq{"id": genreIDs}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	if _, err := l.db.Exec(ctx, sql, args); err != nil {
		return err
	}

	return nil
}
