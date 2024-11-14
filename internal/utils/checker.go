package utils

import "errors"

func SaveError(buff, err error) error {
	if err != nil {
		return errors.Join(buff, err)
	}
	return buff
}
