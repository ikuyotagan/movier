package movies

import (
	"context"

	"github.com/ikuyotagan/movier/internal/models"
)

func (m *Manager) GetPersons(ctx context.Context) ([]*models.Person, error) {
	persons, err := m.db.Persons().All(ctx, nil)
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (m *Manager) CreatePersons(ctx context.Context, persons []*models.Person) ([]uint64, error) {
	ids, err := m.db.Persons().Create(ctx, persons)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
