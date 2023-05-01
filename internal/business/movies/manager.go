package movies

import "github.com/ikuyotagan/movier/internal/database"

// Manager ...
type Manager struct {
	db database.Database
}

// NewManager ...
func NewManager(db database.Database) *Manager {
	return &Manager{db: db}
}
