package database

import "github.com/pkg/errors"

var (
	// ErrEmptyStruct for checking nil structures
	ErrEmptyStruct = errors.New("empty structure")
)
