package task1

type StrengthPotion struct {
	warrior IWarrior
}

func (s *StrengthPotion) getDamage() int {
	warriorDamage := s.warrior.getDamage()
	return warriorDamage + 20
}
