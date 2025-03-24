package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExist        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	defination, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defination, nil
}

func (d Dictionary) Add(word, defination string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = defination
	case nil:
		return ErrWordExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, defination string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = defination
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
