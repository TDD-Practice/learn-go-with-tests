package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (value string, err error) {
	value, ok := d[key]

	if !ok {
		return value, errors.New("key does not exist")
	}

	return value, nil
}
