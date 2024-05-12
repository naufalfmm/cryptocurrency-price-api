package sqliteOrm

import (
	"github.com/mattn/go-sqlite3"
)

func IsUniqueConstraintError(err error) bool {
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return true
		}
	}

	return false
}
