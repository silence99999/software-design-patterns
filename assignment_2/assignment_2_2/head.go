package assignment_2_2

type Head struct{}

func (h *Head) makeBall() IBall {
	return &HeadBall{
		Ball{
			logo:     "Head",
			material: "leather",
		},
	}
}

func (h *Head) makeRacket() IRacket {
	return &HeadRacket{
		Racket{
			logo:     "Head",
			material: "ash wood",
		},
	}
}
