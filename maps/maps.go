package maps

type Dictionary map[string]string

const (
	ErrorKeyNotFound       = DictionaryErr("key does not exist")
	ErrorKeyAlreadyExixsts = DictionaryErr("key already exists")
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
	if d[key] != "" {
		return "", ErrorKeyAlreadyExixsts
	}
	d[key] = value
	return d[key], nil
}
