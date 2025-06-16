package dbx

import "errors"

var (
	ErrInvalidConfig     = errors.New("invalid configuration")
	ErrUnsupportedDriver = errors.New("unsupported database driver")
	ErrConnectionFailed  = errors.New("failed to establish database connection")
)
