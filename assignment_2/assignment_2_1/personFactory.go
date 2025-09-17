package assignment_2_1

import "errors"

func getPerson(name string) (IPerson, error) {
	if name == "Erhan" {
		return newErhan(), nil
	} else if name == "Temirlan" {
		return newTemirlan(), nil
	}

	return nil, errors.New("Not found")
}
