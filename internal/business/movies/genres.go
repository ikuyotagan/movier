package movies

import (
	"context"
	"github.com/ikuyotagan/movier/internal/models"
)

func (m *Manager) GetGenres(ctx context.Context) ([]*models.Genre, error) {
	genres, err := m.db.Genres().All(ctx, nil)
	if err != nil {
		return nil, err
	}
	return genres, nil
}

func (m *Manager) CreateGenres(ctx context.Context, genres []*models.Genre) ([]uint64, error) {
	ids, err := m.db.Genres().Create(ctx, genres)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
