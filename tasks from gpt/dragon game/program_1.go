package main

import "fmt"

func calculateRawDamage(baseDamage int, weaponBonus int) int {
	return baseDamage + weaponBonus
}

func applyArmor(damage int, armor int) int {
	armorDamage := damage - armor
	if armorDamage < 1 {
		return 1
	} else {
		return armorDamage
	}
}

func calculateHitDamage(
	baseDamage int,
	weaponBonus int,
	armor int,
	hitNumber int,
) (int, bool) {
	damage := applyArmor(calculateRawDamage(baseDamage, weaponBonus), armor)
	if hitNumber%3 == 0 {
		return damage*2, true
	} else {
		return damage, false
	}
}

func attackDragon(dragonHealth int, damage int) (int, bool) {
	dragonHealth = dragonHealth - damage
	if dragonHealth > 0 {
		return dragonHealth, false
	} else {
		return 0, true
	}
}

func attackHero(heroHealth int, dragonDamage int) (int, bool) {
	heroHealth = heroHealth - dragonDamage
	if heroHealth > 0 {
		return heroHealth, true
	} else {
		return 0, false
	}
}

func heal(heroHealth int, healingPotions int) (int, int) {
	if healingPotions > 0 {
		heroHealth += 20
		healingPotions -=1
		fmt.Println("Ты успешно похилился +20 хп, текущее здоровье", heroHealth, "зелек осталось:", healingPotions)
		return heroHealth, healingPotions
	} else {
		fmt.Println("Не осталось лечебных зелий!")
		return heroHealth, healingPotions
	}
}

func main() {
	var baseDamage int	// базовый урон героя
	var weaponBonus int // бонус оружия
	var heroHealth int // хп героя
	var dragonArmor int // броня дракона
	var dragonHealth int // хп дракона
	var dragonDamage int // урон дракона
	var attacks int // кол-во доступных атак
	var healingPotions int // кол-во зелий

	fmt.Print("Базовый урон героя: ")
	fmt.Scan(&baseDamage)

	fmt.Print("Бонус оружия: ")
	fmt.Scan(&weaponBonus)

	fmt.Print("Броня дракона: ")
	fmt.Scan(&dragonArmor)

	fmt.Print("Здоровье дракона: ")
	fmt.Scan(&dragonHealth)

	fmt.Print("Количество атак: ")
	fmt.Scan(&attacks)

	fmt.Print("Хп чела: ")
	fmt.Scan(&heroHealth)

	fmt.Print("Атака дракона: ")
	fmt.Scan(&dragonDamage)
	
	fmt.Print("Сколько зелий хочешь взять с собой")
	fmt.Scan(&healingPotions)
	
	alive := true
	defeated := false
	for hitNumber := 1; !(hitNumber == attacks + 1); hitNumber++ {
		damage, critical := calculateHitDamage(baseDamage, weaponBonus, dragonArmor, hitNumber)
		dragonHealth, defeated = attackDragon(dragonHealth, damage)
		heroHealth, alive = attackHero(heroHealth, dragonDamage)
		if !alive {
			defeated = false
			break
		}
		if heroHealth <= 5 {
			heroHealth, healingPotions = heal(heroHealth, healingPotions)
		}
		if defeated {
			fmt.Println("Удар", hitNumber, "нанесено", damage, "урона. Здоровье дракона:", dragonHealth)
			fmt.Println("Дракон побеждён за", hitNumber, "ударов!")
			fmt.Println("У тебя осталось", heroHealth, "здоровья")
			break
		} else if !critical {
			fmt.Println("Удар", hitNumber, "нанесено", damage, "урона. Здоровье дракона:", dragonHealth)
			fmt.Println("У тебя осталось", heroHealth, "здоровья")
		} else {
			fmt.Println("Удар", hitNumber, "КРИТИЧЕСКИЙ УДАР! Нанесено", damage, "урона. Здоровье дракона:", dragonHealth)
			fmt.Println("У тебя осталось", heroHealth, "здоровья")
		}
	}
	if !defeated {
		fmt.Println("ты сдох")
	}
}