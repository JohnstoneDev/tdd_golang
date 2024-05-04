package dictionary

import (
	"errors"
)

type Dictionary map[string]string

type DictionaryErr string

// Global error variables
var (
	ErrNotFound = errors.New("could not find the word you're looking for")
	ErrWordExists = errors.New("cannot add a word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

// Formats an error message
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Method to add an item to the Dictionary
func (d Dictionary) Add(word, definition string) (error) {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound :
		d[word] = definition
	case nil :
		return ErrWordExists
	default :
		return nil
	}

	return nil
}

// Method ot update the value of an item in the dictionary
func (d Dictionary) Update(word, newDefinition string) (error) {
	_, err := d.Search(word)

	switch err {
		case ErrNotFound:
			return ErrWordDoesNotExist
		case nil :
			d[word] = newDefinition
		default:
			return err
	}
	return nil
}

// Method to delete the value of an item in the Dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}