package assignment_2_2

type Wilson struct{}

func (w *Wilson) makeRacket() IRacket {
	return &WilsonRacket{
		Racket{
			logo:     "Wilson",
			material: "wood",
		},
	}
}

func (w *Wilson) makeBall() IBall {
	return &WilsonBall{
		Ball{
			logo:     "Wilson",
			material: "rubber",
		},
	}
}
