package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/ikuyotagan/movier/internal/database"
	"github.com/ikuyotagan/movier/internal/models"
)

const (
	personsTable = "persons"
)

// PersonsRepository имплементация
type PersonsRepository struct {
	db QueryExecutor
}

// NewPersonsRepository конструктор
func NewPersonsRepository(db QueryExecutor) database.PersonsRepository {
	return &PersonsRepository{db: db}
}

// Create persons
func (l PersonsRepository) Create(ctx context.Context, persons []*models.Person) ([]uint64, error) {
	if persons == nil {
		return nil, database.ErrEmptyStruct
	}

	//qb := sq.
	//	Insert(personsTable).
	//	Columns(
	//		"code",
	//		"title",
	//	).
	//	Values(
	//		person.Code,
	//		person.Title,
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

// All get persons
func (l PersonsRepository) All(ctx context.Context, filter *models.PersonsFilter) ([]*models.Person, error) {
	//qb := sq.Select("code", "title").
	//	From(personsTable).
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

// All get persons
func (l PersonsRepository) Count(ctx context.Context, filter *models.PersonsFilter) (uint64, error) {
	//qb := sq.Select("code", "title").
	//	From(personsTable).
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

// Update update person
func (l PersonsRepository) Update(ctx context.Context, person *models.Person) (*models.Person, error) {
	if person == nil {
		return nil, database.ErrEmptyStruct
	}

	//qb := sq.
	//	Update(personsTable).
	//	Where(sq.Eq{"code": person.Code}).
	//	Set("title", person.Title).
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

// Delete person
func (l PersonsRepository) Delete(ctx context.Context, personIDs []uint64) error {
	qb := sq.Delete(personsTable).
		Where(sq.Eq{"id": personIDs}).
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
