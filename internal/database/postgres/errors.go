package postgres

import "github.com/pkg/errors"

var (
	// ErrUnknownExecutor ...
	ErrUnknownExecutor = errors.New("unknown postgres executor")
	// ErrEmptyScheduleIDs ...
	ErrEmptyScheduleIDs = errors.New("empty schedule ids array")
)
