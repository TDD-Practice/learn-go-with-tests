package maps

type Dictionary map[string]string

const (
	ErrorKeyNotFound       = DictionaryErr("key does not exist")
	ErrorKeyAlreadyExixsts = DictionaryErr("key already exists")
	ErrorKeyDoesNotExixst  = DictionaryErr("cannot update an non-existing key")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (value string, err error) {
	value, ok := d[key]

	if !ok {
		return value, ErrorKeyNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) (new_value string, err error) {
	_, searchErr := d.Search(key)

	switch searchErr {
	case ErrorKeyNotFound:
		d[key] = value
		return d[key], nil
	case nil:
		return "", ErrorKeyAlreadyExixsts
	default:
		return "", searchErr
	}
}

func (d Dictionary) Update(key, value string) (new_value string, err error) {
	_, searchErr := d.Search(key)

	switch searchErr {
	case ErrorKeyNotFound:
		return "", ErrorKeyDoesNotExixst
	case nil:
		d[key] = value
		return d[key], nil
	default:
		return "", searchErr
	}
}

func (d Dictionary) Delete(key string) (delete_key string, err error) {
	_, searchErr := d.Search(key)

	switch searchErr {
	case ErrorKeyNotFound:
		return "", ErrorKeyDoesNotExixst
	case nil:
		delete(d, key)
		return key, nil
	default:
		return "", searchErr
	}
}
