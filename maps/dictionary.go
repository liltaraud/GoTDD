package dictionary

const (
	ErrWordNotFound = DictionaryErr("No entry for such key")
	ErrWordExists   = DictionaryErr("Word already exists")
	ErrUpdateFailed = DictionaryErr("The word you're trying to update doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	if _, ok := d[word]; !ok {
		return ErrUpdateFailed
	}

	d[word] = newDefinition
	return nil
}

func (d Dictionary) Add(word string, definition string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
