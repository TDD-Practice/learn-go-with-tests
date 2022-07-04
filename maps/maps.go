package maps

import "errors"

type Dictionary map[string]string

var ErrorKeyNotFound = errors.New("key does not exist")
var ErrorKeyAlreadyExixsts = errors.New("key already exists")

func (d Dictionary) Search(key string) (value string, err error) {
	value, ok := d[key]

	if !ok {
		return value, ErrorKeyNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) (new_value string, err error) {
	d[key] = value
	return d[key], nil
}
