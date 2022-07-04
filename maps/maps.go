package maps

import "errors"

type Dictionary map[string]string

var ErrorKeyNotFound = errors.New("key does not exist")

func (d Dictionary) Search(key string) (value string, err error) {
	value, ok := d[key]

	if !ok {
		return value, ErrorKeyNotFound
	}

	return value, nil
}
