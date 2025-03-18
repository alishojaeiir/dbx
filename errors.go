package dbx

import "errors"

var (
	ErrUnsupportedDriver = errors.New("unsupported database driver")
	ErrConnectionFailed  = errors.New("failed to establish database connection")
)
