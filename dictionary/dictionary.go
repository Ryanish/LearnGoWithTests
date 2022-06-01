package main

type Dictionary map[string]string

type DictionaryErr string

// method reminder - a method is just a function with a receiver argument
// Search "method" here has a receiver type of Dictionary named "d"
// We then return two items - the string and the error
func (d Dictionary) Search(word string) (string, error) {
	// definition is assigned the var for the dictionary map key
	// "ok" is a boolean true/false value that reports whether the value succeeded
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	// Go can use "nil" without first declaring it
	return definition, nil
}

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dictionary map[string]string, word string) string {
	return ""
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
