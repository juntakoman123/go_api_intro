package expense

import "errors"

type Name string

const maxNameLength = 15

func newName(name string) (Name, error) {
	if len(name) > maxNameLength {
		return "", errors.New("expense name must be 15 characters or less")
	}

	return Name(name), nil
}
