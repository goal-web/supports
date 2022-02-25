package utils

import "github.com/pkg/errors"

func NoPanic(fn func()) (err error) {
	defer func() {
		if value := recover(); value != nil {
			switch v := value.(type) {
			case error:
				err = v
			default:
				err = errors.Errorf("%v", v)
			}
		}
	}()

	fn()

	return
}
