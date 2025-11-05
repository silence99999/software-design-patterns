package task1

import "fmt"

func Print() {
	warriorHands := &Hands{}

	//warriorWithSword := &Sword{
	//	warrior: warriorHands,
	//}



	warriorWithSwordAndStrengthPotion := &StrengthPotion{
		warrior: &Sword{
			warrior: &Hands{},
		},
	}

	warriorWithStrengthPotion := &StrengthPotion{
		warrior: warriorHands,
	}

	fmt.Printf("Damage of warrior with sword and Strength potion is %d\n", warriorWithSwordAndStrengthPotion.getDamage())
	fmt.Printf("Damage of warrior with strength potion %d\n", warriorWithStrengthPotion.getDamage())
}
