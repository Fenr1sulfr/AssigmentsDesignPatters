package main

import (
	"fmt"
)

// Races
const (
	RaceBloodElf = "Blood Elf"
	RaceDraenei  = "Draenei"
	RaceDwarf    = "Dwarf"
	RaceGnome    = "Gnome"
	RaceHuman    = "Human"
	RaceNightElf = "Night Elf"
	RaceOrc      = "Orc"
	RaceTauren   = "Tauren"
	RaceTroll    = "Troll"
	RaceUndead   = "Undead"
)

// Classes
const (
	ClassPriest      = "Priest"
	ClassMage        = "Mage"
	ClassWarlock     = "Warlock"
	ClassRogue       = "Rogue"
	ClassFeral       = "Feral"
	ClassBoomkin     = "Boomkin"
	ClassHunter      = "Hunter"
	ClassShaman      = "Shaman"
	ClassWarrior     = "Warrior"
	ClassPaladin     = "Paladin"
	ClassDeathKnight = "Death Knight"
)

// Armor types
const (
	Cloth     = "Cloth"
	Leather   = "Leather"
	ChainMail = "Chain Mail"
	Plate     = "Plate"
)

// CharacterFactory is a singleton that manages character creation.
type CharacterFactory struct{}

func (cf *CharacterFactory) CreateCharacter(race, class, name string) Character {
	// You can extend this to support more races and classes.
	raceFactory := cf.GetRaceFactory(race)
	return raceFactory.CreateCharacter(name, class)
}

func (cf *CharacterFactory) GetRaceFactory(race string) RaceFactory {
	// You can extend this to support more races and factories.
	switch race {
	case RaceHuman:
		return &HumanFactory{}
	case RaceOrc:
		return &OrcFactory{}
	// Add more cases for other races
	default:
		return nil
	}
}

// RaceFactory is the abstract factory interface for creating characters of different races.
type RaceFactory interface {
	CreateCharacter(name, class string) Character
}

// ClassStrategy is the strategy interface for equipping armor based on a character's class.
type ClassStrategy interface {
	EquipArmor(character Character)
	MainStat() string
}

// Character is the base character interface.
type Character interface {
	GetRace() string
	GetClass() string
	GetName() string
	EquipArmor()
	EquipSetArmor(armorType string) bool
	ApplyDebuff()
}

// CharacterDecorator is a decorator pattern for adding abilities to characters.
type CharacterDecorator struct {
	Character
	Ability string
}

func (cd *CharacterDecorator) EquipArmor() {
	cd.Character.EquipArmor()
	fmt.Printf("Adding %s ability to %s's armor.\n", cd.Ability, cd.GetName())
}

func (cd *CharacterDecorator) ApplyDebuff() {
	cd.Character.ApplyDebuff()
	fmt.Printf("%s suffers a debuff because of wearing non-class armor.\n", cd.GetName())
}

// HumanFactory is a concrete factory for creating Human characters.
type HumanFactory struct{}

func (hf *HumanFactory) CreateCharacter(name, class string) Character {
	return &HumanCharacter{
		Race:  RaceHuman,
		Class: class,
		Name:  name,
	}
}

// OrcFactory is a concrete factory for creating Orc characters.
type OrcFactory struct{}

func (of *OrcFactory) CreateCharacter(name, class string) Character {
	return &OrcCharacter{
		Race:  RaceOrc,
		Class: class,
		Name:  name,
	}
}

// DwarfFactory is a concrete factory for creating Dwarf characters.
type DwarfFactory struct{}

func (df *DwarfFactory) CreateCharacter(name, class string) Character {
	return &DwarfCharacter{
		Race:  RaceDwarf,
		Class: class,
		Name:  name,
	}
}

// ElfFactory is a concrete factory for creating Elf characters.
type BloodElfFactory struct{}

func (ef *BloodElfFactory) CreateCharacter(name, class string) Character {
	return &BloodElfCharacter{
		Race:  RaceBloodElf,
		Class: class,
		Name:  name,
	}
}

// UndeadFactory is a concrete factory for creating Undead characters.
type UndeadFactory struct{}

func (uf *UndeadFactory) CreateCharacter(name, class string) Character {
	return &UndeadCharacter{
		Race:  RaceUndead,
		Class: class,
		Name:  name,
	}
}

// GnomeFactory is a concrete factory for creating Gnome characters.
type GnomeFactory struct{}

func (gf *GnomeFactory) CreateCharacter(name, class string) Character {
	return &GnomeCharacter{
		Race:  RaceGnome,
		Class: class,
		Name:  name,
	}
}

// TrollFactory is a concrete factory for creating Troll characters.
type TrollFactory struct{}

func (tf *TrollFactory) CreateCharacter(name, class string) Character {
	return &TrollCharacter{
		Race:  RaceTroll,
		Class: class,
		Name:  name,
	}
}

// NightElfFactory is a concrete factory for creating Night Elf characters.
type NightElfFactory struct{}

func (nef *NightElfFactory) CreateCharacter(name, class string) Character {
	return &NightElfCharacter{
		Race:  RaceNightElf,
		Class: class,
		Name:  name,
	}
}

// TaurenFactory is a concrete factory for creating Tauren characters.
type TaurenFactory struct{}

func (tf *TaurenFactory) CreateCharacter(name, class string) Character {
	return &TaurenCharacter{
		Race:  RaceTauren,
		Class: class,
		Name:  name,
	}
}

// Implement concrete character types for each race and class.
// Ensure that each character type applies debuffs when wearing non-class armor.

// Example for a Warrior character:
type WarriorCharacter struct {
	Race  string
	Class string
	Name  string
}

func (wc *WarriorCharacter) EquipArmor() {
	fmt.Printf("%s equips Warrior armor.\n", wc.GetName())
}

func (wc *WarriorCharacter) EquipSetArmor(armorType string) bool {
	if armorType == "Plate" {
		return true // Warriors can equip Plate armor without debuffs.
	} else {
		return false // Applying debuff for wearing non-Plate armor.
	}
}

func (wc *WarriorCharacter) ApplyDebuff() {
	fmt.Printf("%s suffers a Strength debuff.\n", wc.GetName())
}

// Implement similar character types for other classes (e.g., RogueCharacter, MageCharacter, etc.) and ensure they apply debuffs as needed.

// ...

// Implement the concrete character types for each race and class (e.g., HumanCharacter, OrcCharacter, etc.).
// Ensure that these characters apply debuffs for wearing non-class armor as per your requirements.

func main() {
	// Singleton: Create a character creation system (CharacterFactory) as a single instance.
	characterFactory := &CharacterFactory{}

	// Let the player choose the race and class for customization.
	race := RaceHuman
	class := ClassWarrior
	name := "Timur"

	// Create the character with the chosen race and class.
	character := characterFactory.CreateCharacter(race, class, name)

	// Strategy: Equip armor based on the character's class.
	armorStrategy := getArmorStrategy(class)
	character.EquipArmorStrategy(armorStrategy)

	// Decorator: Add abilities to characters.
	enhancedCharacter := &CharacterDecorator{
		Character: character,
		Ability:   "Strength",
	}
	enhancedCharacter.EquipArmor()
	enhancedCharacter.ApplyDebuff()
}

func getArmorStrategy(class string) ClassStrategy {
	// Implement strategies for each class based on the main stat.
	switch class {
	case ClassWarrior, ClassPaladin, ClassDeathKnight:
		return &WarriorStrategy{}
	case ClassRogue, ClassFeral, ClassBoomkin:
		return &RogueStrategy{}
	case ClassHunter:
		return &HunterStrategy{}
	case ClassShaman:
		return &ShamanStrategy{}
	case ClassPriest, ClassMage, ClassWarlock, ClassBoomkin:
		return &CasterStrategy{}
	// Add more cases for other classes
	default:
		return &DefaultStrategy{}
	}
}

// Implement concrete armor strategies for each class based on main stat.

// Implement methods to check if a character can equip a specific armor set.
