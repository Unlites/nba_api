package utils

import (
	"errors"
	"strconv"
)

func ParseId(queryParam string) (int64, error) {
	idInput, err := strconv.ParseInt(queryParam, 10, 64)
	if err != nil {
		return 0, err
	}

	if idInput < 1 {
		return 0, errors.New("not positive id")
	}

	return idInput, nil
}
