package database

import (
	"context"

	"github.com/ikuyotagan/movier/internal/models"
)

type (
	// Transaction выполняет транзакционные операции посредством временно живущего объекта базы данных
	Transaction func(db Database) error

	// Database ...
	Database interface {
		Movies() MoviesRepository
		Persons() PersonsRepository
		Genres() GenresRepository
		Transaction(ctx context.Context, transaction Transaction) error
	}

	// MoviesRepository для работы с фильмами
	MoviesRepository interface {
		Create(ctx context.Context, movie []*models.Movie) ([]*models.Movie, error)
		All(ctx context.Context, filter *models.MoviesFilter) ([]*models.Movie, error)
		Count(ctx context.Context, filter *models.MoviesFilter) (uint64, error)
		Update(ctx context.Context, movie *models.Movie) (*models.Movie, error)
		Delete(ctx context.Context, movieIDs []uint64) error
	}

	// PersonsRepository ...
	PersonsRepository interface {
		Create(ctx context.Context, persons []*models.Person) ([]*models.Person, error)
		All(ctx context.Context, filter *models.PersonsFilter) ([]*models.Person, error)
		Count(ctx context.Context, filter *models.PersonsFilter) (uint64, error)
		Update(ctx context.Context, person *models.Person) (*models.Person, error)
		Delete(ctx context.Context, personIDs []uint64) error
	}

	// GenresRepository ...
	GenresRepository interface {
		Create(ctx context.Context, genres []*models.Genre) ([]*models.Genre, error)
		All(ctx context.Context, filter *models.GenresFilter) ([]*models.Genre, error)
		Count(ctx context.Context, filter *models.GenresFilter) (uint64, error)
		Update(ctx context.Context, rule *models.Genre) (*models.Genre, error)
		Delete(ctx context.Context, genreIDs []uint64) error
	}
)
