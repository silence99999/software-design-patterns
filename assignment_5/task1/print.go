package task1

import "fmt"

func Print() {
	warriorHands := &Hands{}

	warriorWithSword := &Sword{
		warrior: warriorHands,
	}

	warriorWithSwordAndStrengthPotion := &StrengthPotion{
		warrior: warriorWithSword,
	}

	fmt.Printf("Damage of warrior with sword and Strength potion is %d\n", warriorWithSwordAndStrengthPotion.getDamage())
}
