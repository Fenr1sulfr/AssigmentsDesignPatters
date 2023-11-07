package main

import "fmt"

type IWeapon interface {
	setName(name string)
	setDamage(damage int)
	getName() string
	getDamage() int
	getAttack() string
}

type Weapon struct {
	name   string
	damage int
}

func (w *Weapon) setName(name string) {
	w.name = name
}
func (w *Weapon) setDamage(damage int) {
	w.damage = damage
}
func (w *Weapon) getName() string {
	return w.name
}
func (w *Weapon) getDamage() int {
	return w.damage
}
func (w *Weapon) getAttack() string {
	return fmt.Sprintf("You're attacking with %b", w.getDamage())
}

type Staff struct {
	Weapon
}

func newStaff() IWeapon {
	return &Staff{
		Weapon: Weapon{
			name:   "Wizzard Staff",
			damage: 10,
		},
	}
}

type OneHandSword struct {
	Weapon
}

func newSword() IWeapon {
	return &OneHandSword{
		Weapon: Weapon{
			name:   "One-handed Sword",
			damage: 15,
		},
	}
}

type TwoHandSword struct {
	Weapon
}

//нужно посмотреть как сделать здесь decorator
func newBigSword() IWeapon {
	return &TwoHandSword{
		Weapon: Weapon{
			name:   "Two-handed Sword",
			damage: 45,
		},
	}
}

type Bow struct {
	Weapon
}

func newBow() IWeapon {
	return &Bow{
		Weapon: Weapon{
			name:   "Bow",
			damage: 20,
		},
	}
}

type Stick struct {
	Weapon
}

func newStick() IWeapon {
	return &Stick{
		Weapon: Weapon{
			name:   "Magic Stick",
			damage: 8,
		},
	}
}

type Shield struct {
	Weapon
}

func newShield() IWeapon {
	return &Shield{
		Weapon: Weapon{
			name:   "Shield",
			damage: 0,
		},
	}
}

type BlockedWeapon struct {
	Weapon
}

func newBlockedWeapon() IWeapon {
	return &BlockedWeapon{
		Weapon: Weapon{
			name:   "You can't use this hand",
			damage: 0,
		},
	}
}
func choiceMelee(class string) (IWeapon, IWeapon, error) {
	var input int
	fmt.Println("What would you like Tank or Damage Dealer?")
	fmt.Println("Send 1 - Tank, 2 - Damage Dealer")
	fmt.Scanln(&input)
	if input == 1 {
		return newShield(), newSword(), nil
	} else {
		return newBigSword(), newBlockedWeapon(), nil
	}
}
func choiceDistance(class string) (IWeapon, IWeapon, error) {
	var input int
	fmt.Println("What would you like Stick or Staff?")
	fmt.Println("Send 1 - Stick, 2 - Staff")
	fmt.Scanln(&input)
	if input == 1 {
		return newStick(), newBlockedWeapon(), nil
	} else {
		return newStaff(), newBlockedWeapon(), nil
	}
}

//Factor
func getWeapon(class string) (IWeapon, IWeapon, error) {
	//swords
	if class == "Warrior" || class == "Paladin" || class == "Death Knight" {
		return choiceMelee(class)
	}
	if class == "Mage" || class == "Priest" {
		return choiceDistance(class)
	}
	return newBlockedWeapon(), newBlockedWeapon(), fmt.Errorf("Unknown class")
}
