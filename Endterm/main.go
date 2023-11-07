package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	"Priest",
	"Mage",
	"Warlock",
	"Rogue",
	"Druid",
	"Hunter",
	"Shaman",
	"Warrior",
	"Paladin",
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
	1:  "Priest",
	2:  "Mage",
	3:  "Warlock",
	4:  "Rogue",
	5:  "Druid",
	6:  "Hunter",
	7:  "Shaman",
	8:  "Warrior",
	9:  "Paladin",
	10: "Death Knight",
}

type Character struct {
	Name  string
	Race  string
	Class string
}

type CharacterBuilder struct {
	character *Character
}

func NewCharacterBuilder() *CharacterBuilder {
	return &CharacterBuilder{character: &Character{}}
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

func (cb *CharacterBuilder) Build() *Character {
	return cb.character
}

func main() {
	// Using the Builder pattern
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
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		race, ok := raceOptions[raceIndex]
		if !ok {
			fmt.Println("Please enter a valid number from the list.")
		} else {
			selectedRace = race
			break
		}
	}

	var selectedClass string
	for {
		fmt.Println("Choose your character's class:")
		for i, validClass := range validClasses {
			fmt.Printf("%d. %s\n", i+1, validClass)
		}

		classNumber := getUserInput()
		classIndex, err := strconv.Atoi(classNumber)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		class, ok := classOptions[classIndex]
		if !ok {
			fmt.Println("Please enter a valid number from the list.")
		} else {
			selectedClass = class
			break
		}
	}

	character := NewCharacterBuilder().
		SetName(name).
		SetRace(selectedRace).
		SetClass(selectedClass).
		Build()

	// Display character attributes
	fmt.Printf("Character Name: %s\n", character.Name)
	fmt.Printf("Character Race: %s\n", character.Race)
	fmt.Printf("Character Class: %s\n", character.Class)
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// Remove trailing newline character
	input = strings.TrimSpace(input)
	return input
}
