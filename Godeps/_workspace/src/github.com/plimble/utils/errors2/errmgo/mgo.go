package errmgo

import (
	"github.com/plimble/utils/errors2"
)

func Err(err error) error {
	if err == nil {
		return nil
	}

	if err.Error() == "not found" {
		return errors2.NewNotFound("not found")
	}

	return errors2.NewInternal(err.Error())
}
