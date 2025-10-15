package task1

type Sword struct {
	warrior IWarrior
}

func (s *Sword) getDamage() int {
	warriorDamage := s.warrior.getDamage()
	return warriorDamage + 10
}
