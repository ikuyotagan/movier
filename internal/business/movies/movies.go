package movies

import (
	"context"
	"github.com/ikuyotagan/movier/internal/models"
	"golang.org/x/sync/errgroup"
)

func (m *Manager) GetMovies(ctx context.Context, filter *models.MoviesFilter) (movies []*models.Movie, err error) {
	group, errCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		movies, err = m.db.Movies().All(errCtx, filter)
		return err
	})
	genresMap := make(map[uint64]*models.Genre)
	group.Go(func() error {
		var genres []*models.Genre
		genres, err = m.GetGenres(errCtx)
		if err != nil {
			return err
		}

		for _, genre := range genres {
			genresMap[genre.Id] = genre
		}
		return nil
	})
	if err = group.Wait(); err != nil {
		return nil, err
	}

	for _, movie := range movies {
		genres := make([]*models.Genre, 0, len(movie.Genres))
		for _, genre := range movie.Genres {
			genres = append(genres, genresMap[genre.Id])
		}
		movie.Genres = genres
	}
	return movies, nil
}

func (m *Manager) CreateMovies(ctx context.Context, moviesCreateParams []*models.MovieCreateParams) ([]uint64, error) {
	movies := make([]*models.Movie, 0, len(moviesCreateParams))
	for _, createParams := range moviesCreateParams {
		genres := make([]*models.Genre, 0, len(createParams.GenreIds))
		for _, genre := range createParams.GenreIds {
			genres = append(genres, &models.Genre{
				Id: genre,
			})
		}
		movies = append(movies, &models.Movie{
			// ToDo:: дописать остальные поля
			ExternalInfo: createParams.ExternalInfo,
			Genres:       genres,
			DateOfFilmed: createParams.DateOfReleased,
			PosterUrl:    createParams.PosterUrl,
			Name:         createParams.Name,
			Rating:       createParams.Rating,
			//Country:      createParams.Country,
		})
	}
	ids, err := m.db.Movies().Create(ctx, movies)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (m *Manager) DeleteMovies(ctx context.Context, ids []uint64) error {
	return m.db.Movies().Delete(ctx, ids)
}
