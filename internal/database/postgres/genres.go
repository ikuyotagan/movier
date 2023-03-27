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
func (l GenresRepository) Create(ctx context.Context, genre []*models.Genre) ([]*models.Genre, error) {
	if len(genre) == 0 {
		return nil, database.ErrEmptyStruct
	}
	//
	//qb := sq.
	//	Insert(genresTable).
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

// All get genres
func (l GenresRepository) All(ctx context.Context, filter *models.GenresFilter) ([]*models.Genre, error) {
	//qb := sq.Select("code", "title").
	//	From(genresTable).
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

// All get genres
func (l GenresRepository) Count(ctx context.Context, filter *models.GenresFilter) (uint64, error) {
	//qb := sq.Select("code", "title").
	//	From(genresTable).
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
func (l GenresRepository) Update(ctx context.Context, layer *models.Genre) (*models.Genre, error) {
	//if layer == nil {
	//	return database.ErrEmptyStruct
	//}
	//
	//qb := sq.
	//	Update(genresTable).
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
func (l GenresRepository) Delete(ctx context.Context, genreIDs []uint64) error {
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
