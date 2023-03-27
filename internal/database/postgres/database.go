package postgres

import (
	"context"
	"sync"

	"github.com/ikuyotagan/movier/internal/database"
)

// Database имплементация
type Database struct {
	db QueryExecutor

	movies     database.MoviesRepository
	initMovies sync.Once

	persons     database.PersonsRepository
	initPersons sync.Once

	genres     database.GenresRepository
	initGenres sync.Once
}

// NewDatabase конструктор
func NewDatabase(db QueryExecutor) database.Database {
	return &Database{db: db}
}

// Transaction ...
func (d *Database) Transaction(ctx context.Context, transactionFn database.Transaction) error {
	tx, err := d.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	transactionalDB := &Database{db: tx}
	if err := transactionFn(transactionalDB); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// Movies ...
func (d *Database) Movies() database.MoviesRepository {
	d.initMovies.Do(func() {
		d.movies = NewMoviesRepository(d.db)
	})
	return d.movies
}

// Persons ...
func (d *Database) Persons() database.PersonsRepository {
	d.initPersons.Do(func() {
		d.persons = NewPersonsRepository(d.db)
	})
	return d.persons
}

// Genres ...
func (d *Database) Genres() database.GenresRepository {
	d.initGenres.Do(func() {
		d.genres = NewGenresRepository(d.db)
	})
	return d.genres
}
