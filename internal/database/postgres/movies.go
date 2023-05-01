package postgres

import (
	"context"
	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/ikuyotagan/movier/internal/database"
	"github.com/ikuyotagan/movier/internal/models"
)

const (
	moviesTable       = "movies"
	movieGenreTable   = "movie_genre"
	externalInfoTable = "external_info"
	moviePersonTable  = "movie_person"
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
func (l *MoviesRepository) Create(ctx context.Context, movies []*models.Movie) ([]uint64, error) {
	if movies == nil {
		return nil, database.ErrEmptyStruct
	}
	qb := sq.
		Insert(moviesTable).
		Columns(
			"name",
			"date_of_release",
			"poster_url",
			"rating",
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)
	for _, movie := range movies {
		qb = qb.Values(
			movie.Name,
			movie.DateOfFilmed.AsTime(),
			movie.PosterUrl,
			movie.Rating,
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

	moviesGenres := make(map[uint64][]uint64, len(movies))
	moviesExternalInfo := make(map[uint64][]*models.ExternalInfo, len(movies))
	group, groupContext := errgroup.WithContext(ctx)
	for i, movie := range movies {
		for _, genre := range movie.Genres {
			moviesGenres[ids[i]] = append(moviesGenres[ids[i]], genre.Id)
		}
		moviesExternalInfo[ids[i]] = movie.ExternalInfo
	}

	group.Go(func() error {
		if err := l.createMovieGenresLinks(groupContext, moviesGenres); err != nil {
			return err
		}
		return nil
	})

	group.Go(func() error {
		if err := l.createMovieExternalInfo(groupContext, moviesExternalInfo); err != nil {
			return err
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		return nil, err
	}

	return ids, nil
}

// All get movies
func (l *MoviesRepository) All(ctx context.Context, filter *models.MoviesFilter) ([]*models.Movie, error) {
	limit := filter.Limit
	if limit == 0 {
		limit = 50
	}
	offset := uint64(0)
	if filter.Page != 0 {
		offset = (filter.Page - 1) * limit
	}
	qb := sq.Select(
		"id",
		"name",
		"date_of_release",
		"poster_url",
		"rating",
		"array_agg("+movieGenreTable+".genre_id)",
		"max("+externalInfoTable+".ext_id)",
		"max("+externalInfoTable+".ext_source)",
		"max("+externalInfoTable+".ext_rating)",
		"max("+externalInfoTable+".rates)",
	).
		From(moviesTable).
		Join(movieGenreTable + " on " + moviesTable + ".id = " + movieGenreTable + ".movie_id").
		Join(externalInfoTable + " on " + moviesTable + ".id = " + externalInfoTable + ".movie_id").
		OrderBy("id").
		Limit(limit).
		Offset(offset).
		GroupBy(moviesTable + ".id").
		PlaceholderFormat(sq.Dollar)
	if len(filter.Ids) != 0 {
		qb = qb.Where(sq.Eq{"id": filter.Ids})
	}
	if len(filter.GenreIds) != 0 {
		qb = qb.Where(sq.Eq{movieGenreTable + ".genre_id": filter.GenreIds})
	}

	result := make([]*models.Movie, 0)
	sql, args, err := qb.ToSql()
	log.Info(ctx, sql)
	if err != nil {
		return nil, err
	}

	rows, err := l.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		movie := models.Movie{}
		extInfo := models.ExternalInfo{}
		var dateOfFilmed time.Time
		var genreIDs []uint64
		var rating string
		var extID string
		var extSource string

		err := rows.Scan(
			&movie.Id,
			&movie.Name,
			&dateOfFilmed,
			&movie.PosterUrl,
			&rating,
			&genreIDs,
			&extID,
			&extSource,
			&extInfo.Rating,
			&extInfo.Rates,
		)
		intExtID, _ := strconv.Atoi(extID)
		extInfo.ExtId = int64(intExtID)
		extInfo.Source = movier.ExternalSource(movier.ExternalSource_value[extSource])
		movie.Rating = movier.MPAARating(movier.MPAARating_value[rating])
		movie.DateOfFilmed = timestamppb.New(dateOfFilmed)
		if err != nil {
			return nil, err
		}
		movie.ExternalInfo = []*models.ExternalInfo{&extInfo}

		for _, id := range genreIDs {
			movie.Genres = append(movie.Genres, &models.Genre{
				Id: id,
			})
		}
		result = append(result, &movie)
	}
	return result, nil
}

// All get movies
func (l *MoviesRepository) Count(ctx context.Context, filter *models.MoviesFilter) (uint64, error) {
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
func (l *MoviesRepository) Update(ctx context.Context, layer *models.Movie) (*models.Movie, error) {
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
func (l *MoviesRepository) Delete(ctx context.Context, movieIDs []uint64) error {
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

func (l *MoviesRepository) createMovieGenresLinks(ctx context.Context, movieGenreIDs map[uint64][]uint64) error {
	if len(movieGenreIDs) == 0 {
		return database.ErrEmptyStruct
	}
	qb := sq.
		Insert(movieGenreTable).
		Columns(
			"movie_id",
			"genre_id",
		).
		PlaceholderFormat(sq.Dollar)
	for movie, genres := range movieGenreIDs {
		for _, genre := range genres {
			qb = qb.Values(
				movie,
				genre,
			)
		}
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	_, err = l.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (l *MoviesRepository) createMovieExternalInfo(ctx context.Context, moviesInfo map[uint64][]*models.ExternalInfo) error {
	if len(moviesInfo) == 0 {
		return database.ErrEmptyStruct
	}
	qb := sq.
		Insert(externalInfoTable).
		Columns(
			"movie_id",
			"ext_id",
			"ext_source",
			"ext_rating",
			"rates",
		).
		PlaceholderFormat(sq.Dollar)
	for movie, infos := range moviesInfo {
		for _, info := range infos {
			qb = qb.Values(
				movie,
				strconv.Itoa(int(info.ExtId)),
				info.Source,
				info.Rating,
				info.Rates,
			)
		}
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	_, err = l.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
