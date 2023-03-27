package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/ikuyotagan/movier/internal/database"
	"github.com/ikuyotagan/movier/internal/models"
)

const (
	moviesTable = "movies"
)

// MoviesRepository имплементация
type MoviesRepository struct {
	db QueryExecutor
}

// NewMoviesRepository конструктор
func NewMoviesRepository(db QueryExecutor) database.MoviesRepository {
	return &MoviesRepository{db: db}
}

// Create movies
func (l MoviesRepository) Create(ctx context.Context, movies []*models.Movie) ([]*models.Movie, error) {
	if movies == nil {
		return nil, database.ErrEmptyStruct
	}

	//qb := sq.
	//	Insert(moviesTable).
	//	Columns(
	//		"code",
	//		"title",
	//	).
	//	Values(
	//		layer.Code,
	//		layer.Title,
	//	).
	//	PlaceholderFormat(sq.Dollar)
	//
	//executor, err := executor(ctx, l.db, role.Write)
	//if err != nil {
	//	return err
	//}
	//
	//if _, err := executor.Execx(ctx, qb); err != nil {
	//	return err
	//}

	return nil, nil
}

// All get movies
func (l MoviesRepository) All(ctx context.Context, filter *models.MoviesFilter) ([]*models.Movie, error) {
	//qb := sq.Select("code", "title").
	//	From(moviesTable).
	//	PlaceholderFormat(sq.Dollar)
	//
	//if filter != nil {
	//	if len(filter.Codes) > 0 {
	//		qb = qb.Where(sq.Eq{"code": filter.Codes})
	//	}
	//
	//	if len([]rune(filter.Query)) > 3 {
	//		qb = qb.Where("title ilike ?", "%"+filter.Query+"%")
	//	}
	//}
	//
	//qb = qb.OrderBy("code")
	//
	//executor, err := executor(ctx, l.db, role.Read)
	//if err != nil {
	//	return nil, err
	//}
	//
	//result := make([]*models.Layer, 0)
	//if err := executor.Selectx(ctx, &result, qb); err != nil {
	//	return nil, err
	//}
	return nil, nil
}

// All get movies
func (l MoviesRepository) Count(ctx context.Context, filter *models.MoviesFilter) (uint64, error) {
	//qb := sq.Select("code", "title").
	//	From(moviesTable).
	//	PlaceholderFormat(sq.Dollar)
	//
	//if filter != nil {
	//	if len(filter.Codes) > 0 {
	//		qb = qb.Where(sq.Eq{"code": filter.Codes})
	//	}
	//
	//	if len([]rune(filter.Query)) > 3 {
	//		qb = qb.Where("title ilike ?", "%"+filter.Query+"%")
	//	}
	//}
	//
	//qb = qb.OrderBy("code")
	//
	//executor, err := executor(ctx, l.db, role.Read)
	//if err != nil {
	//	return nil, err
	//}
	//
	//result := make([]*models.Layer, 0)
	//if err := executor.Selectx(ctx, &result, qb); err != nil {
	//	return nil, err
	//}
	return 0, nil
}

// Update update layer
func (l MoviesRepository) Update(ctx context.Context, layer *models.Movie) (*models.Movie, error) {
	if layer == nil {
		return nil, database.ErrEmptyStruct
	}

	//qb := sq.
	//	Update(moviesTable).
	//	Where(sq.Eq{"code": layer.Code}).
	//	Set("title", layer.Title).
	//	PlaceholderFormat(sq.Dollar)
	//
	//sql, args, err := qb.ToSql()
	//if err != nil {
	//	return err
	//}
	//
	//if _, err := l.db.Exec(ctx, sql, args); err != nil {
	//	return err
	//}

	return nil, nil
}

// Delete delete layer
func (l MoviesRepository) Delete(ctx context.Context, movieIDs []uint64) error {
	qb := sq.Delete(moviesTable).
		Where(sq.Eq{"id": movieIDs}).
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
