package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

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

type Wand struct {
	Weapon
}

func newWand() IWeapon {
	return &Wand{
		Weapon: Weapon{
			name:   "Magic Wand",
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
	fmt.Println("What would you like Wand or Staff?")
	fmt.Println("Send 1 - Wand, 2 - Staff")
	fmt.Scanln(&input)
	if input == 1 {
		return newWand(), newBlockedWeapon(), nil
	} else {
		return newStaff(), newBlockedWeapon(), nil
	}
}

// Factor
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

var validRaces = []string{
	"Blood Elf",
	"Draenei",
	"Dwarf",
	"Gnome",
	"Human",
	"Night Elf",
	"Orc",
	"Tauren",
	"Troll",
	"Undead",
}

var validClasses = []string{
	"Paladin",
	"Mage",
	"Warrior",
	"Priest",
	"Death Knight",
}

var raceOptions = map[int]string{
	1:  "Blood Elf",
	2:  "Draenei",
	3:  "Dwarf",
	4:  "Gnome",
	5:  "Human",
	6:  "Night Elf",
	7:  "Orc",
	8:  "Tauren",
	9:  "Troll",
	10: "Undead",
}

var classOptions = map[int]string{
	1: "Paladin",
	2: "Mage",
	3: "Warrior",
	4: "Priest",
	5: "Death Knight",
}

type Character struct {
	Name    string
	Race    string
	Class   string
	Weapon  IWeapon
	Offhand IWeapon
}

type CharacterBuilder struct {
	character *Character
}

var builderInstance *CharacterBuilder
var builderOnce sync.Once

func GetCharacterBuilder() *CharacterBuilder {
	builderOnce.Do(func() {
		builderInstance = &CharacterBuilder{character: &Character{}}
	})
	return builderInstance
}

func NewCharacterBuilder() *CharacterBuilder {
	return GetCharacterBuilder()
}

func (cb *CharacterBuilder) SetName(name string) *CharacterBuilder {
	cb.character.Name = name
	return cb
}

func (cb *CharacterBuilder) SetRace(race string) *CharacterBuilder {
	cb.character.Race = race
	return cb
}

func (cb *CharacterBuilder) SetClass(class string) *CharacterBuilder {
	cb.character.Class = class
	return cb
}
func (cb *CharacterBuilder) SetWeapon(weapon IWeapon) *CharacterBuilder {
	cb.character.Weapon = weapon
	return cb
}

func (cb *CharacterBuilder) SetOffhand(weapon IWeapon) *CharacterBuilder {
	cb.character.Offhand = weapon
	return cb
}
func (cb *CharacterBuilder) Build() *Character {
	return cb.character
}

type NzothWeaponDecor struct {
	weapon IWeapon
}

func (nd *NzothWeaponDecor) setName(name string) {
	nd.weapon.setName(name + ", Curse of N'Zoth")
}
func (nd *NzothWeaponDecor) setDamage(damage int) {
	nd.weapon.setDamage(damage + 10)
}
func (nd *NzothWeaponDecor) getName() string {
	return nd.weapon.getName()
}
func (nd *NzothWeaponDecor) getDamage() int {
	return nd.weapon.getDamage()
}
func (nd *NzothWeaponDecor) getAttack() string {
	return nd.weapon.getAttack() + " , N'Zoth helps you, but hurts your mind"
}

type AzerothDecor struct {
	weapon IWeapon
}

func (az *AzerothDecor) setName(name string) {
	az.weapon.setName(name + ", Consecrated by Azeroth")
}
func (az *AzerothDecor) setDamage(damage int) {
	az.weapon.setDamage(damage + 8)
}
func (az *AzerothDecor) getName() string {
	return az.weapon.getName()
}
func (az *AzerothDecor) getDamage() int {
	return az.weapon.getDamage()
}
func (az *AzerothDecor) getAttack() string {
	return az.weapon.getAttack() + " , Azeroth gives you more power"
}
func (az *AzerothDecor) AzWeap(name string, damage int) {
	az.setName(name)
	az.setDamage(damage)
	// attackDescription := az.getAttack()
	// //fmt.Printf("Weapon: %s\nDamage: %d\nAttack: %s\n", az.getName(), az.getDamage(), attackDescription)
}
func (nd *NzothWeaponDecor) NWeap(name string, damage int) {
	nd.setName(name)
	nd.setDamage(damage)
	// attackDescrind.getAttack()
	// // fmt.Printf("Weapon: %s\nDamage: %d\nAttack: %s\n", nd.getName(), nd.getDamage(), attackDescription)
}

func main() {
	fmt.Print("Do you want to create a new character (N) or load an existing character (L): ")
	choice := getUserInput()

	var character *Character

	if strings.ToLower(choice) == "n" {
		fmt.Print("Enter your character's name: ")
		name := getUserInput()

		var selectedRace string
		for {
			fmt.Println("Choose your character's race:")
			for i, validRace := range validRaces {
				fmt.Printf("%d. %s\n", i+1, validRace)
			}

			raceNumber := getUserInput()
			raceIndex, err := strconv.Atoi(raceNumber)
			if err != nil || raceIndex < 1 || raceIndex > len(raceOptions) {
				fmt.Println("Please enter a valid number.")
				continue
			}

			selectedRace = raceOptions[raceIndex]
			break
		}

		var selectedClass string
		for {
			fmt.Println("Choose your character's class:")
			for i, validClass := range validClasses {
				fmt.Printf("%d. %s\n", i+1, validClass)
			}

			classNumber := getUserInput()
			classIndex, err := strconv.Atoi(classNumber)
			if err != nil || classIndex < 1 || classIndex > len(classOptions) {
				fmt.Println("Please enter a valid number.")
				continue
			}

			selectedClass = classOptions[classIndex]
			break
		}

		var selectedWeapon IWeapon
		var selectedOffhand IWeapon
		selectedWeapon, selectedOffhand, err := getWeapon(selectedClass)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Print("Are you an honorable hero striving to protect Azeroth, or do you embrace the darkness and walk the path of a cunning villain? (A/N)\n")
		side := getUserInput()
		if strings.ToLower(side) == "n" {
			nzothWeapon := &NzothWeaponDecor{weapon: selectedWeapon}
			nzothWeapon.NWeap(nzothWeapon.getName(), selectedWeapon.getDamage())
			ofNzothWeapon := &NzothWeaponDecor{weapon: selectedOffhand}
			if selectedOffhand.getDamage() > 0 {
				ofNzothWeapon.NWeap(ofNzothWeapon.getName(), selectedOffhand.getDamage())
				selectedOffhand = ofNzothWeapon
			}
			selectedWeapon = nzothWeapon
		}
		if strings.ToLower(side) == "a" {
			azerothWeapon := &AzerothDecor{weapon: selectedWeapon}
			azerothWeapon.AzWeap(azerothWeapon.getName(), selectedWeapon.getDamage())
			selectedWeapon = azerothWeapon
			ofAzerothWeapon := &AzerothDecor{weapon: selectedOffhand}
			if selectedOffhand.getDamage() > 0 {
				ofAzerothWeapon.AzWeap(ofAzerothWeapon.getName(), selectedOffhand.getDamage())
				selectedOffhand = ofAzerothWeapon
			}
		}
		character = NewCharacterBuilder().
			SetName(name).
			SetRace(selectedRace).
			SetClass(selectedClass).
			SetWeapon(selectedWeapon).
			SetOffhand(selectedOffhand).
			Build()

	}

	fmt.Printf("Character Name: %s\n", character.Name)
	fmt.Printf("Character Race: %s\n", character.Race)
	fmt.Printf("Character Class: %s\n", character.Class)
	fmt.Printf("Character Weapon: %s\n", character.Weapon.getName())
	fmt.Printf("Character Offhand: %s\n", character.Offhand.getName())
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}
