package errsqlx

import (
	"github.com/plimble/utils/errors2"
)

func Err(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "sql: no rows in result set":
		return errors2.NewNotFound("not found")
	}

	return errors2.NewInternal(err.Error())
}
