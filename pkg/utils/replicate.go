package utils

import (
	"time"
)

func DoWithTries(fn func() error, attemps int, delley time.Duration) (err error) {
	for attemps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delley)
			attemps--
			continue
		}
		return nil
	}

	return
}
