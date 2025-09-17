package assignment_2_2

import (
	"errors"
)

type ITennisFactory interface {
	makeRacket() IRacket
	makeBall() IBall
}

func GetTennisFactory(brand string) (ITennisFactory, error) {
	if brand == "Wilson" {
		return &Wilson{}, nil
	} else if brand == "Head" {
		return &Head{}, nil
	}

	return nil, errors.New("Not Found")
}
